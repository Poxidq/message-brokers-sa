package main

import (
	"log"

	"filter/config"
	"filter/rabbitmq"
	"filter/service"
)

func main() {
	config.LoadConfig()
	conn, ch, err := rabbitmq.InitRabbitMQ("FilterQueue")
	if err != nil {
		log.Fatalf("Error initializing RabbitMQ: %s", err)
	}
	defer conn.Close()
	defer ch.Close()

	if err := rabbitmq.ConsumeMessage(ch, "FilterQueue", service.FilterMessage); err != nil {
		log.Fatalf("Error consuming message: %s", err)
	}
}