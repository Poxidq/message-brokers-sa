package main

import (
	"log"

	"screaming/config"
	"screaming/rabbitmq"
	"screaming/service"
)

func main() {
	config.LoadConfig()
	conn, ch, err := rabbitmq.InitRabbitMQ("ScreamingQueue")
	if err != nil {
		log.Fatalf("Error initializing RabbitMQ: %s", err)
	}
	defer conn.Close()
	defer ch.Close()

	if err := rabbitmq.ConsumeMessage(ch, "ScreamingQueue", service.MakeUppercase); err != nil {
		log.Fatalf("Error consuming message: %s", err)
	}
}