package utils

import (
	"encoding/json"
	"event-driven/config"
	"event-driven/models"
	"fmt"

	"github.com/streadway/amqp"
)

type Message struct {
	Message string `json:"message"`
	User    string `json:"user"`
}

func InitRabbitMQ(queueName string) (*amqp.Connection, *amqp.Channel, error) {
	// RabbitMQ connection parameters
	rabbitMQHost := config.CFG.RabbitMQHost
	rabbitMQPort := config.CFG.RabbitMQPort
	rabbitMQUser := config.CFG.RabbitMQUser
	rabbitMQPassword := config.CFG.RabbitMQPassword

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

// PublishMessage sends a Message to the specified queue
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
