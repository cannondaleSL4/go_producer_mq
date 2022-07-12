package rabbitMq

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"go_producer_mq/config"
	. "go_producer_mq/data"
	"log"
	"os"
)

type RabbitMsg struct {
	QueueName string     `json:"queueName"`
	Reply     UsersOrder `json:"reply"`
}

// channel to publish rabbit messages
var rchan = make(chan RabbitMsg, 10)

func InitProducer(cfg config.Config) {
	// conn
	conn, err := NewRabbitMQConn(cfg)
	if err != nil {
		log.Printf("ERROR: fail init consumer: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("INFO: done init producer conn")

	// create channel
	amqpChannel, err := conn.Channel()
	if err != nil {
		log.Printf("ERROR: fail create channel: %s", err.Error())
		os.Exit(1)
	}

	for {
		select {
		case msg := <-rchan:
			// marshal
			data, err := proto.Marshal(&msg.Reply)
			if err != nil {
				log.Printf("ERROR: fail marshal: %s", err.Error())
				continue
			}

			// publish message
			err = amqpChannel.Publish(
				"",            // exchange
				msg.QueueName, // routing key
				false,         // mandatory
				false,         // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        data,
				},
			)
			if err != nil {
				log.Printf("ERROR: fail publish msg: %s", err.Error())
				continue
			}

			log.Printf("INFO: published msg: %v to: %s", msg.Reply, msg.QueueName)
		}
	}
}

func NewRabbitMQConn(cfg config.Config) (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	return amqp.Dial(connAddr)
}
