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

type rabbitMsg struct {
	QueueName string     `json:"queueName"`
	Order     UsersOrder `json:"reply"`
}

var rchan = make(chan rabbitMsg, 10)

func InitProducer(cfg config.Config) {
	// conn
	conn, err := newRabbitMQConn(cfg)
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
			data, err := proto.Marshal(&msg.Order)
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

			log.Printf("INFO: published msg: %v to: %s", msg.Order, msg.QueueName)
		}
	}
}

func newRabbitMQConn(cfg config.Config) (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	return amqp.Dial(connAddr)
}

func PublishMessage(message *UsersOrder) {
	msg := rabbitMsg{
		QueueName: "storage",
		Order:     *message,
	}
	rchan <- msg
}
