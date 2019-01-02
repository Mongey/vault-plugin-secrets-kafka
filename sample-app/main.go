package main

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

const topic = "my-topic"

func main() {
	bootstrapServers := []string{"localhost:9092"}
	cfg, err := newTLSConfig("client.cert", "private.key", "ca.cert")
	if err != nil {
		log.Fatalf("%s", err)
	}
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Version = sarama.V2_0_0_0
	kafkaConfig.ClientID = "terraform-provider-kafka"
	kafkaConfig.Net.TLS.Enable = true
	kafkaConfig.Net.TLS.Config = cfg
	sarama.Logger = log.New(os.Stderr, "", log.LstdFlags)

	consumeAllMessages(topic, bootstrapServers, kafkaConfig)
	produce(topic, bootstrapServers, kafkaConfig)
}
