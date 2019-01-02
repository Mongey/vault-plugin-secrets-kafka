package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func produce(topic string, bootstrapServers []string, kafkaConfig *sarama.Config) {
	kafkaConfig.Producer.Return.Successes = true
	p, err := sarama.NewSyncProducer(bootstrapServers, kafkaConfig)

	if err != nil {
		log.Fatalf("%s", err)
	}

	key := "test"
	value := "hello"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}
	_, _, err = p.SendMessage(msg)

	if err != nil {
		log.Fatalf("%s", err)
	}
}
