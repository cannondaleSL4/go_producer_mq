package service

import . "go_producer_mq/data"

type RabbitMsgSent struct {
	QueueName string     `json:"queueName"`
	Message   UsersOrder `json:"message"`
}

var pchan = make(chan RabbitMsgSent)

func PublishMessageTemp(message *UsersOrder) {

	msg := RabbitMsgSent{
		QueueName: "storage",
		Message:   *message,
	}
	pchan <- msg
}
