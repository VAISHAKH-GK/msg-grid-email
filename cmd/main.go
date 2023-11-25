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
	godotenv.Load(".env")
	var amqpUrl = os.Getenv("AMQP_URL")

	var connectRabbitMq,err = amqp.Dial(amqpUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer connectRabbitMq.Close()

	channel,err := connectRabbitMq.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	channel.QueueDeclare(
		"send-email",
		true,
		true,
		false,
		false,
		nil,
	)
	
	messages,err := channel.Consume(
		"send-email",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for msg := range messages {
			var m map[string]interface{}
			json.Unmarshal(msg.Body,&m)
			fmt.Println(m)
		}
	}
}
