package consumer

import (
    "encoding/json"
    "fmt"
    "notification-service/internal/events"
    "os"

    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type (
    NotifyHandler interface {
        CreateHandler(data *events.NotificationEvent) error
    }

    WorkerNotifyHandler interface {
        CreateWorkerHandler(data *events.NotificationEvent) error
    }
)

type Consumer struct {
    consumer            *kafka.Consumer
    notifyHandler       NotifyHandler
    workerNotifyHandler WorkerNotifyHandler
}

func NewConsumer(
    notifyHandler NotifyHandler,
    workerNotifyHandler WorkerNotifyHandler) *Consumer {

    config := &kafka.ConfigMap{
        "bootstrap.servers": "kafka:29092",
        "group.id":          "myGroup",
        "auto.offset.reset": "smallest",
    }
    consumer, err := kafka.NewConsumer(config)
    if err != nil {
        _ = fmt.Errorf("error new consumer %e", err)
    }
    topic := "async-topic"
    err = consumer.SubscribeTopics([]string{topic}, nil)
    if err != nil {
        panic(err)
    }
    return &Consumer{
        consumer:            consumer,
        notifyHandler:       notifyHandler,
        workerNotifyHandler: workerNotifyHandler,
    }
}

func (c *Consumer) Start() {

    for {
        ev := c.consumer.Poll(100)

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

            fmt.Println(event, "Event from Consumer.go!")

            err = c.notifyHandler.CreateHandler(&event)
            if err != nil {
                fmt.Printf("handler create error %v", err)
                break
            }

            err = c.workerNotifyHandler.CreateWorkerHandler(&event)
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
    c.consumer.Close()
}
