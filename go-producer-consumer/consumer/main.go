package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "a-sample-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{"game-events"}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topics: %s", err)
	}

	fmt.Println("Consumer started...")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Consumed message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)]n", err, msg)
		}
	}
}
