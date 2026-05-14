package producer

import (
	"fmt"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewProducer() {

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
	})

	if err != nil {
		panic(err)
	}
	defer producer.Close()
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for event := range producer.Events() {

			switch ev := event.(type) {

			case *kafka.Message:

				// delivery report
				if ev.TopicPartition.Error != nil {
					fmt.Println("delivery failed:", ev.TopicPartition.Error)
				} else {
					fmt.Printf(
						"message delivered to %s [%d] at offset %v\n",
						*ev.TopicPartition.Topic,
						ev.TopicPartition.Partition,
						ev.TopicPartition.Offset,
					)
				}
			}
		}
	}()

	topic := "test-topic"

	// отправляем сообщения
	for i := 1; i <= 10; i++ {

		message := fmt.Sprintf("hello kafka %d", i)

		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(message),
		}, nil)

		if err != nil {
			fmt.Println("produce error:", err)
			continue
		}

		fmt.Println("sent:", message)

		time.Sleep(1 * time.Second)
	}

	wg.Wait()
	// ждем отправки сообщений
	producer.Flush(5000)

	fmt.Println("producer finished")

}
