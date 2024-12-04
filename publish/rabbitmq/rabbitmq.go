package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"publish/config"
	"publish/models"

	"github.com/streadway/amqp"
)

func InitRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	// RabbitMQ connection parameters
	rabbitMQHost := config.CFG.RabbitMQHost
	rabbitMQPort := config.CFG.RabbitMQPort
	rabbitMQUser := config.CFG.RabbitMQUser
	rabbitMQPassword := config.CFG.RabbitMQPassword
	log.Printf("connecting to amqp://%s:%s@%s:%s/", rabbitMQUser, rabbitMQPassword, rabbitMQHost, rabbitMQPort)
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitMQUser, rabbitMQPassword, rabbitMQHost, rabbitMQPort))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	queueName := "PublishQueue"
	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to declare queue %s: %v", queueName, err)
	}

	return conn, ch, nil
}

type SendEmail func(string, string) error

func ConsumeMessage(ch *amqp.Channel, queueName string, callback SendEmail) error {
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to consume messages: %s", err)
	}

	log.Println("Publish service is listening for messages...")
	for msg := range msgs {
		log.Printf("Received message: %s", msg.Body)

		var message models.Message
		err := json.Unmarshal(msg.Body, &message)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		err = callback("New Message", message.Message)
		if err != nil {
			log.Printf("Error sending email: %v", err)
			return err
		}
	}

	return nil
}
