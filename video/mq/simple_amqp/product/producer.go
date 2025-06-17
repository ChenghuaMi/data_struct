package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}

// connect
// channel
// queue
// publish
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672") //connect
	failOnError(err, "failed to connect to mq")
	defer conn.Close()
	ch, err := conn.Channel() //channel
	failOnError(err, "failed to open a channel")
	defer ch.Close()
	//queue
	q, err := ch.QueueDeclare("hello", true, false, false, false, nil)
	failOnError(err, "failed to declare on queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := "hello msg"
	//publish
	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	failOnError(err, "failed to publish a msg")
	log.Printf("send %s \n", body)
}
