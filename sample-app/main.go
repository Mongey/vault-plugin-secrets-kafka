package main

import (
	"io/ioutil"
	"log"

	"github.com/Mongey/terraform-provider-kafka/kafka"
	"github.com/Shopify/sarama"
)

const topic = "my-topic"

func main() {
	_client, err := client("localhost:9092", "client.cert", "private.key", "ca.cert")
	if err != nil {
		log.Fatalf("%s", err)
	}

	consumeAllMessages(topic, bootstrapServers, kafkaConfig)
	produce(topic, bootstrapServers, kafkaConfig)
}

func client(broker, caLocation, clientCertLocation, clientKeyLocation string) (*sarama.Client, error) {
	brokers := []string{broker}
	caCert, err := ioutil.ReadFile(caLocation)
	if err != nil {
		return nil, err
	}

	clientCert, err := ioutil.ReadFile(clientCertLocation)
	if err != nil {
		return nil, err
	}

	clientKey, err := ioutil.ReadFile(clientKeyLocation)
	if err != nil {
		return nil, err
	}

	config := &kafka.Config{
		BootstrapServers: &brokers,
		CACert:           string(caCert),
		ClientCert:       string(clientCert),
		ClientCertKey:    string(clientKey),
		SkipTLSVerify:    false,
		TLSEnabled:       true,
		Timeout:          100,
	}

	client, err := kafka.NewClient(config)
	if err != nil {
		return nil, err
	}

	c := client.SaramaClient()

	return &c, nil
}
