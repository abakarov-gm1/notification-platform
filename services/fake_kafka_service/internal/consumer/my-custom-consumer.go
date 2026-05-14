package consumer

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewsConsumer() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "myGroup",
		"auto.offset.reset": "smallest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(fmt.Sprintf("Failed to create consumer: %v", err))
	}
	topic := "async-topic"

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer initialized")

	for {
		ev := consumer.Poll(100)

		if ev == nil {
			continue
		}

		switch e := ev.(type) {

		case *kafka.Message:

			fmt.Printf(
				"Message: %s\nTopic: %s\nPartition: %d\nOffset: %v\n\n",
				string(e.Value),
				*e.TopicPartition.Topic,
				e.TopicPartition.Partition,
				e.TopicPartition.Offset,
			)

		case kafka.Error:

			fmt.Fprintf(os.Stderr, "Kafka error: %v\n", e)

		default:

			fmt.Printf("Ignored event: %T\n", e)
		}
	}
	consumer.Close()
}
