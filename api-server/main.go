package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const QUEUE_NAME = "payments"

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Panicf("%s: %s", msg, err)
	}
}

func main() {
	logrus.Info("Application is starting....123")

	err := LoadConfig()
	failOnError(err, "Failed to create config")

	db, err := setupDB(AppConfig)
	failOnError(err, "Failed to create db connection")

	messageQueue, err := setupMessageQueue(AppConfig)
	failOnError(err, "Failed to setup message queue")

	router := gin.Default()
	RegisterRoutes(router, messageQueue, db)
	router.Run(":8888")
}

func setupDB(cfg Config) (*gorm.DB, error) {
	return NewDbConnection(cfg)
}

func setupMessageQueue(cfg Config) (*RabbitMQConnection, error) {
	return NewRabbitMQConnection(cfg, QUEUE_NAME)
}
