package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "test_mi",
		Partition: 0,
		GroupID:   "test-group-1",
	})
	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("topic: %s,key:%s,value:%s\n", m.Topic, m.Key, m.Value)
		if err = r.CommitMessages(ctx, m); err != nil {
			panic(fmt.Sprintf("fail to commit error:%s", err))
		}
	}
}
