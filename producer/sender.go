package main

import (
	"context"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

var channel amqp091.Channel

func Connect(connectionUrl string) {
	connection, err := amqp091.Dial(connectionUrl)
	failOnError(err, "Error while establishing connection to AMQP client")
	defer connection.Close()

	channel, err := connection.Channel()
	failOnError(err, "Error while establishing connection to channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Error while establishing queue")

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
