package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"hello", // Queue name
		false,   // Durable
		false,   // Delete when unused
		false,   // Exclusive
		false,   // No-wait
		nil,     // Arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // Queue
		"",     // Consumer
		true,   // Auto Ack
		false,  // Exclusive
		false,  // No Local
		false,  // No Wait
		nil,    // Args
	)
	failOnError(err, "Failed to register a consumer")

	// Create a channel to receive signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Wait for messages
	go func() {
		for msg := range msgs {
			log.Printf(" [x] Received %s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-signals
	log.Printf(" [*] Exiting...")
}
