package main

import (
	"bytes"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func fail(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	fail(err, "connect error")
	ch, err := conn.Channel()
	fail(err, "channel error")
	err = ch.ExchangeDeclare("logs", amqp.ExchangeFanout, true, false, false, false, nil)
	fail(err, "fanout error")
	q, err := ch.QueueDeclare("", true, false, false, false, nil)
	fail(err, "failed to queue")
	err = ch.QueueBind(q.Name, "", "logs", false, nil)
	//err = ch.Qos(1, 0, false)
	fail(err, "bind error")
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	fail(err, "consumer error")
	wait := make(chan struct{})
	go func() {
		for msg := range msgs {
			count := bytes.Count(msg.Body, []byte("."))
			time.Sleep(time.Duration(count) * time.Second)
			fmt.Println(fmt.Sprintf("%s", msg.Body))
			msg.Ack(false)
		}
	}()
	<-wait
}
