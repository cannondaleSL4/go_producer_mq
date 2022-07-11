package rabbitMq

import (
	"fmt"
	"go_producer_mq/config"
)

func NewRabbitMQConn(cfg *config.Config) (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	return amqp.Dial(connAddr)
}
