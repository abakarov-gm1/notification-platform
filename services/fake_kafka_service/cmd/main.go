package main

import "fake/internal/producer"

func main() {
	p := producer.NewProducer()
	topics := []string{"async-topic"}

	produce := p.Init()
	p.Send(topics[0], produce)

}
