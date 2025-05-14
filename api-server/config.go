package main

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost      string `envconfig:"db_host" required:"true"`
	DBUser      string `envconfig:"db_user" required:"true"`
	DBPassword  string `envconfig:"db_password" required:"true"`
	DBName      string `envconfig:"db_name" required:"true"`
	RabbitMQUrl string `envconfig:"rabbit_mq_url" required:"true"`
}

var AppConfig Config

func LoadConfig() error {
	if err := envconfig.Process("", &AppConfig); err != nil {
		return err
	}

	return nil
}
