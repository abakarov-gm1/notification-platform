package producer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewsProducer() {

	configs := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
	}

	producer, err := kafka.NewProducer(configs)
	if err != nil {
		_ = fmt.Errorf("ошибкка с консюмером  %e", err)
	}
	defer producer.Close()

	fmt.Println("Init producer...")

	topic := "notification"

	for _, word := range []string{"hello", "from", "test", "connect Kafka"} {

		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}

		err := producer.Produce(message, nil)

		if err != nil {
			fmt.Println("produce error:", err)
			continue
		}

	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition.Error)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

}
