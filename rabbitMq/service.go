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

type producerStruct struct {
	conf config.Config
}

var rchan = make(chan rabbitMsg, 10)

func (p *producerStruct) InitProducer(cfg config.Config) {

	p.conf = cfg

	conn, err := p.newRabbitMQConn()
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

func (p *producerStruct) newRabbitMQConn() (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		p.conf.RabbitMQ.User,
		p.conf.RabbitMQ.Password,
		p.conf.RabbitMQ.Host,
		p.conf.RabbitMQ.Port,
	)
	return amqp.Dial(connAddr)
}

func (p *producerStruct) PublishMessage(message *UsersOrder) {
	msg := rabbitMsg{
		QueueName: p.conf.RabbitMQ.Queue,
		Order:     *message,
	}
	rchan <- msg
}
