package config

import (
	"fmt"

	env "github.com/caarlos0/env/v11"
)

type config struct {
	RabbitMQUser     string `env:"RABBITMQ_USER"`
	RabbitMQPort     string `env:"RABBITMQ_PORT"`
	RabbitMQPassword string `env:"RABBITMQ_PASSWORD"`
	RabbitMQHost     string `env:"RABBITMQ_HOST"`
}

var CFG config

func LoadConfig() {
	err := env.Parse(&CFG)

	res, err := env.ParseAs[config]()
	if err != nil {
		panic(fmt.Sprintf("Error when loading environment %v", err))
	}

	CFG = res
}
