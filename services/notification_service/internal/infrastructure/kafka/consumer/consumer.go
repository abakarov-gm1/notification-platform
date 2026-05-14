package consumer

import (
	"encoding/json"
	"fmt"
	"notification-service/internal/infrastructure/kafka/consumer/events"
	"notification-service/internal/infrastructure/kafka/consumer/handlers"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer struct {
	conf *kafka.ConfigMap
}

func NewConsumer(conf *kafka.ConfigMap) *Consumer {
	return &Consumer{conf: conf}
}

func (c *Consumer) ConsumerCreate() *kafka.Consumer {
	consumer, err := kafka.NewConsumer(c.conf)
	if err != nil {
		_ = fmt.Errorf("error new consumer %e", err)
	}
	topic := "notification"
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		panic(err)
	}
	return consumer
}

func (c *Consumer) Start(consumer *kafka.Consumer, notifHandler *handlers.NotificationHandler) {

	for {
		ev := consumer.Poll(100)

		if ev == nil {
			continue
		}

		switch e := ev.(type) {

		case *kafka.Message:

			var event events.NotificationEvent

			err := json.Unmarshal(e.Value, &event)
			if err != nil {
				fmt.Printf("unmarshal error: %v\n", err)
				break
			}

			err = notifHandler.CreateHandler(&event)
			if err != nil {
				fmt.Printf("handler create error %v", err)
				break
			}

			err = notifHandler.CreateWorkersHandler(&event)
			if err != nil {
				fmt.Printf("handler worker error %v", err)
				break
			}

		case kafka.Error:

			fmt.Fprintf(os.Stderr, "Kafka error: %v\n", e)

		default:

			fmt.Printf("Ignored event: %T\n", e)
		}
	}
	consumer.Close()
}
