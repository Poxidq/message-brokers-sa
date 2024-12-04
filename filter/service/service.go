package service

import (
	"log"
	"strings"

	"filter/models"
	"filter/rabbitmq"

	"github.com/streadway/amqp"
)

var stopWords = []string{"bird-watching", "ailurophobia", "mango"}

func FilterMessage(ch *amqp.Channel, message models.Message) error {
	for _, word := range stopWords {
		if strings.Contains(message.Message, word) {
			log.Printf("filtered because message contains: '%s'", word)
			return nil
		}
	}

	log.Printf("passed message: %s", message.Message)

	return rabbitmq.PublishMessage(ch, "ScreamingQueue", message)
}