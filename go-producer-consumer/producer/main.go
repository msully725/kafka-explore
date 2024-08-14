package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer, err := kafka.NewProducer(
		&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	topic := "game-events"

	for {
		message := fmt.Sprintf("Event: %d", rand.Int())
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny},
			Value: []byte(message),
		}, nil)

		if err != nil {
			fmt.Printf("Failed to produce message: %v\n", err)
		}

		fmt.Printf("Produced message: %s\n", message)

		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	}
}
