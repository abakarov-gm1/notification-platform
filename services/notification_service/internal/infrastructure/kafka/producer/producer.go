package producer

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
    producer *kafka.Producer
}

func NewProducer() *Producer {
    configs := &kafka.ConfigMap{
        "bootstrap.servers": "kafka:29092",
    }
    p, err := kafka.NewProducer(configs)

    if err != nil {
        _ = fmt.Errorf("error producer %e", err)
    }
    producer := &Producer{producer: p}

    go producer.HandleEvents()

    return producer
}

func (p *Producer) HandleEvents() {
    for e := range p.producer.Events() {
        switch ev := e.(type) {

        case *kafka.Message:
            if ev.TopicPartition.Error != nil {
                fmt.Printf(
                    "Failed to deliver message: %v\n",
                    ev.TopicPartition.Error,
                )
            } else {
                fmt.Printf(
                    "Delivered to topic %s [%d] offset %v\n",
                    *ev.TopicPartition.Topic,
                    ev.TopicPartition.Partition,
                    ev.TopicPartition.Offset,
                )
            }
        }
    }
}

func (p *Producer) Send(topic string, value []byte) error {
    message := &kafka.Message{
        TopicPartition: kafka.TopicPartition{
            Topic:     &topic,
            Partition: kafka.PartitionAny,
        },
        Value: value,
    }
    return p.producer.Produce(message, nil)
}

func (p *Producer) Close() {
    p.producer.Flush(5000)
    p.producer.Close()
}
