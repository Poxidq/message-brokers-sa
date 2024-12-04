package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"filter/config"
	"filter/models"

	"github.com/streadway/amqp"
)

func InitRabbitMQ(queueName string) (*amqp.Connection, *amqp.Channel, error) {
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

func PublishMessage(ch *amqp.Channel, queueName string, msg models.Message) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}

type FilterMessage func(*amqp.Channel, models.Message) error

func ConsumeMessage(ch *amqp.Channel, queueName string, callback FilterMessage) error {
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

	log.Println("Filter service is listening for messages...")
	for msg := range msgs {
		log.Printf("Received message: %s", msg.Body)

		var message models.Message
		err := json.Unmarshal(msg.Body, &message)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		err = callback(ch, message)
		if err != nil {
			log.Printf("Error filtering message: %v", err)
			return err
		}
	}

	return nil
}
