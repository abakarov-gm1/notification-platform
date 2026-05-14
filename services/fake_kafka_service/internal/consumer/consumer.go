package consumer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewCons() {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",

		// группа consumer
		"group.id": "test-group",

		// читать сначала
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	topic := "test-topic"

	// подписка на topic
	err = consumer.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("consumer started...")

	for {

		msg, err := consumer.ReadMessage(-1)

		if err != nil {
			fmt.Println("consumer error:", err)
			continue
		}

		fmt.Printf(
			"My Consumer! received message: %s | partition: %d | offset: %d\n",
			string(msg.Value),
			msg.TopicPartition.Partition,
			msg.TopicPartition.Offset,
		)
	}
}
