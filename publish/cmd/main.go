package main

import (
	"log"

	"publish/config"
	"publish/rabbitmq"
	"publish/service"
)

func main() {
	config.LoadConfig()
	conn, ch, err := rabbitmq.InitRabbitMQ()
	if err != nil {
		log.Fatalf("Error initializing RabbitMQ: %s", err)
	}
	defer conn.Close()
	defer ch.Close()

	if err := rabbitmq.ConsumeMessage(ch, "PublishQueue", service.SendEmail); err != nil {
		log.Fatalf("Error consuming message: %s", err)
	}
}
