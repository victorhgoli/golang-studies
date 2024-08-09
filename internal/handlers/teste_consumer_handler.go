package handlers

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func HandleMessage(message *kafka.Message) {
	log.Printf("Received message: %s", string(message.Value))
	


	
}
