package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		AllowAutoTopicCreation: true,
		Topic:                  "test_mi",
		Balancer:               &kafka.LeastBytes{},
	}
	defer w.Close()
	err := w.WriteMessages(context.Background(), kafka.Message{
		Key:       []byte("key-a1"),
		Value:     []byte("key-a-value"),
		Partition: 0,
	}, kafka.Message{
		Key:       []byte("key-b1"),
		Value:     []byte("key-b-value"),
		Partition: 0,
	})
	if err != nil {
		panic(fmt.Sprintf("write error:%s", err))
	}
}
