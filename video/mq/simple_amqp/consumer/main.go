package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}

// connect
// channel
// queue
// consumer

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "failed to connect to mq")
	defer conn.Close()
	ch, err := conn.Channel() //channel
	failOnError(err, "failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare( // queue
		"hello",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare a queue")
	// consumer
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "failed to register a consumer")
	var forever = make(chan struct{})
	go func() {
		for d := range msgs {
			log.Printf("received a message: %s", d.Body)
		}
	}()
	<-forever
}
