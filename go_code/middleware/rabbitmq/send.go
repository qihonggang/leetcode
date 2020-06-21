package main

import (
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@127.0.0.1:5672/admin")
	failOnError(err, "Failed to connect to RabbitMQ ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel ")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}
