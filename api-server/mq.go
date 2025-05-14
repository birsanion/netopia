package main

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConnection struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

func NewRabbitMQConnection(cfg Config, name string) (*RabbitMQConnection, error) {
	conn, err := amqp.Dial(cfg.RabbitMQUrl)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	queue, err := ch.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &RabbitMQConnection{
		Conn:    conn,
		Channel: ch,
		Queue:   &queue,
	}, nil
}

func (c *RabbitMQConnection) Publish(ctx context.Context, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return c.Channel.PublishWithContext(ctx,
		"",           // exchange
		c.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
}

func (c *RabbitMQConnection) IsAvailable() bool {
	if c == nil {
		return false
	}

	return !c.Conn.IsClosed()
}
