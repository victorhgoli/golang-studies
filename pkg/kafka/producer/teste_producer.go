package producer

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer interface {
	ProduceMessage() error
}

type producer struct {
}

func NewProducer() (Producer, error) {

	return &producer{}, nil
}

func (p *producer) ProduceMessage() error {

	brokers := "localhost:9092"
	topic := "create-user-topic"
	message := "Hello, Kafka!"

	config := &kafka.ConfigMap{
		"bootstrap.servers": brokers,
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		return err
	}

	deliveryChan := make(chan kafka.Event)

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)

	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		return m.TopicPartition.Error
	} else {
		log.Printf("Delivered message to %v\n", m.TopicPartition)
	}

	close(deliveryChan)
	return nil
}
