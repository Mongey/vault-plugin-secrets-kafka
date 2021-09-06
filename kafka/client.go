package kafka

import (
	"errors"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type TopicMissingError struct {
	msg string
}

func (e TopicMissingError) Error() string { return e.msg }

type void struct{}

var member void

type Client struct {
	client        sarama.Client
	kafkaConfig   *sarama.Config
	config        *Config
	supportedAPIs map[int]int
	topics        map[string]void
}

func NewClient(config *Config) (*Client, error) {
	if config == nil {
		return nil, errors.New("Cannot create client without kafka config")
	}

	if config.BootstrapServers == nil {
		return nil, fmt.Errorf("No bootstrap_servers provided")
	}

	bootstrapServers := *(config.BootstrapServers)
	if bootstrapServers == nil {
		return nil, fmt.Errorf("No bootstrap_servers provided")
	}

	kc, err := config.newKafkaConfig()
	if err != nil {
		log.Printf("[ERROR] Error creating kafka client %v", err)
		return nil, err
	}

	c, err := sarama.NewClient(bootstrapServers, kc)
	if err != nil {
		log.Printf("[ERROR] Error connecting to kafka %s", err)
		return nil, err
	}

	client := &Client{
		client:      c,
		config:      config,
		kafkaConfig: kc,
	}

	err = client.populateAPIVersions()
	if err != nil {
		return client, err
	}

	err = client.extractTopics()

	return client, err
}

func (c *Client) SaramaClient() sarama.Client {
	return c.client
}

func (c *Client) populateAPIVersions() error {
	ch := make(chan []*sarama.ApiVersionsResponseBlock)
	errCh := make(chan error)

	brokers := c.client.Brokers()
	kafkaConfig := c.kafkaConfig
	for _, broker := range brokers {
		go apiVersionsFromBroker(broker, kafkaConfig, ch, errCh)
	}

	clusterApiVersions := make(map[int][2]int) // valid api version intervals across all brokers
	errs := make([]error, 0)
	for i := 0; i < len(brokers); i++ {
		select {
		case brokerApiVersions := <-ch:
			updateClusterApiVersions(&clusterApiVersions, brokerApiVersions)
		case err := <-errCh:
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errors.New(sarama.MultiError{Errors: &errs}.PrettyError())
	}

	c.supportedAPIs = make(map[int]int, len(clusterApiVersions))
	for apiKey, versionMinMax := range clusterApiVersions {
		versionMin := versionMinMax[0]
		versionMax := versionMinMax[1]

		if versionMax >= versionMin {
			c.supportedAPIs[apiKey] = versionMax
		}

		// versionMax will be less than versionMin only when
		// two or more brokers have disjoint version
		// intervals...which means the api is not supported
		// cluster-wide
	}

	return nil
}

func apiVersionsFromBroker(broker *sarama.Broker, config *sarama.Config, ch chan<- []*sarama.ApiVersionsResponseBlock, errCh chan<- error) {
	resp, err := rawApiVersionsRequest(broker, config)

	if err != nil {
		errCh <- err
	} else if resp.Err != sarama.ErrNoError {
		errCh <- errors.New(resp.Err.Error())
	} else {
		ch <- resp.ApiVersions
	}
}

func rawApiVersionsRequest(broker *sarama.Broker, config *sarama.Config) (*sarama.ApiVersionsResponse, error) {
	if err := broker.Open(config); err != nil && err != sarama.ErrAlreadyConnected {
		return nil, err
	}

	defer func() {
		if err := broker.Close(); err != nil && err != sarama.ErrNotConnected {
			log.Fatal(err)
		}
	}()

	return broker.ApiVersions(&sarama.ApiVersionsRequest{})
}

func updateClusterApiVersions(clusterApiVersions *map[int][2]int, brokerApiVersions []*sarama.ApiVersionsResponseBlock) {
	cluster := *clusterApiVersions

	for _, apiBlock := range brokerApiVersions {
		apiKey := int(apiBlock.ApiKey)
		brokerMin := int(apiBlock.MinVersion)
		brokerMax := int(apiBlock.MaxVersion)

		clusterMinMax, exists := cluster[apiKey]
		if !exists {
			cluster[apiKey] = [2]int{brokerMin, brokerMax}
		} else {
			// shrink the cluster interval according to
			// the broker interval

			clusterMin := clusterMinMax[0]
			clusterMax := clusterMinMax[1]

			if brokerMin > clusterMin {
				clusterMinMax[0] = brokerMin
			}

			if brokerMax < clusterMax {
				clusterMinMax[1] = brokerMax
			}

			cluster[apiKey] = clusterMinMax
		}
	}
}

func (c *Client) extractTopics() error {
	topics, err := c.client.Topics()
	if err != nil {
		log.Printf("[ERROR] Error getting topics %s from Kafka", err)
		return err
	}
	log.Printf("[DEBUG] Got %d topics from Kafka", len(topics))
	c.topics = make(map[string]void)
	for _, t := range topics {
		c.topics[t] = member
	}
	return nil
}

func (c *Client) CanAlterReplicationFactor() bool {
	_, ok1 := c.supportedAPIs[45] // https://kafka.apache.org/protocol#The_Messages_AlterPartitionReassignments
	_, ok2 := c.supportedAPIs[46] // https://kafka.apache.org/protocol#The_Messages_ListPartitionReassignments

	return ok1 && ok2
}

func (c *Client) allReplicas() *[]int32 {
	brokers := c.client.Brokers()
	replicas := make([]int32, 0, len(brokers))

	for _, b := range brokers {
		id := b.ID()
		if id != -1 {
			replicas = append(replicas, id)
		}
	}

	return &replicas
}
func (c *Client) versionForKey(apiKey, wantedMaxVersion int) int {
	if maxSupportedVersion, ok := c.supportedAPIs[apiKey]; ok {
		if maxSupportedVersion < wantedMaxVersion {
			return maxSupportedVersion
		}
		return wantedMaxVersion
	}

	return 0
}

func (c *Client) getDescribeAclsRequestAPIVersion() int16 {
	return int16(c.versionForKey(29, 1))
}
func (c *Client) getCreateAclsRequestAPIVersion() int16 {
	return int16(c.versionForKey(30, 1))
}

func (c *Client) getDeleteAclsRequestAPIVersion() int16 {
	return int16(c.versionForKey(31, 1))
}
