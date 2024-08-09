package consumer

import (
	"estudo-test/internal/handlers"
	"estudo-test/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer interface {
	StartTesteConsumer()
}

type consumer struct {
	UserService *service.UserService
	Handler     *func(message *kafka.Message)
}

func NewTesteConsumer(userService *service.UserService) Consumer {
	return &consumer{UserService: userService}
}

func (c *consumer) StartTesteConsumer() {

	brokers := "localhost:9092"
	groupID := "my-group"
	topics := []string{"create-user-topic"}

	config := &kafka.ConfigMap{
		"bootstrap.servers":  brokers,
		"group.id":           groupID,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topics: %v", err)
	}

	listen(consumer, handlers.HandleMessage)
}

func listen(consumer *kafka.Consumer, handlerFunc func(*kafka.Message)) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	run := true

	for run {
		select {
		case sig := <-sigchan:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				log.Printf("Message on %s: %s\n", e.TopicPartition, string(e.Value))
				// Process the message
				//handlers.HandleMessage(e)
				handlerFunc(e)
				consumer.CommitMessage(e)
			case kafka.Error:
				log.Printf("Error: %v\n", e)
				run = false
			default:
				// Ignore other events
			}
		}
	}
}
