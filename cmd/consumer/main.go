package main

import (
	"fmt"

	"github.com/ClaudionorJunior/go-expert-events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "minhafila")
	for msg := range msgs {
		fmt.Println("Received a message: ", string(msg.Body))
		msg.Ack(false)
	}
}
