package service

import (
	"log"
	"strings"

	"screaming/models"
	"screaming/rabbitmq"

	"github.com/streadway/amqp"
)

func MakeUppercase(ch *amqp.Channel, message models.Message) error {
	message.Message = strings.ToUpper(message.Message)
	log.Printf("made uppercase: %s", message.Message)

	return rabbitmq.PublishMessage(ch, "PublishQueue", message)
}