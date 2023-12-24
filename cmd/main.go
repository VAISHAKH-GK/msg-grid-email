package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Load enviornment variables
	godotenv.Load(".env")
	var amqpUrl = os.Getenv("AMQP_URL")

	// connecting to rabbit mq
	var connectRabbitMq, err = amqp.Dial(amqpUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer connectRabbitMq.Close()

	// Connecting rabbit mq channel
	channel, err := connectRabbitMq.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	// Creating rabbit mq channel
	channel.QueueDeclare(
		"send-email",
		true,
		false,
		false,
		false,
		nil,
	)
	channel.Qos(1, 0, false)

	// Consuming messages from rabbit mq
	messages, err := channel.Consume(
		"send-email",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Get messages from rabbiq mq channel
	var forver = make(chan bool)
	GetMessages(messages)
	<-forver
}

func GetMessages(messages <-chan amqp.Delivery) {
	for msg := range messages {
		var m map[string]interface{}
		json.Unmarshal(msg.Body, &m)
		fmt.Println(m)
		msg.Ack(false)
	}
}
