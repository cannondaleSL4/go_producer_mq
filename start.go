package main

import (
	. "go_producer_mq/config"
	. "go_producer_mq/httpserver"
	rabbitMq "go_producer_mq/rabbitMq"
	"log"
)

func main() {
	//data := GetOrder()
	//fmt.Println(data)

	cfgPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	//init rabbit
	//go InitProducer(*cfg)
	go rabbitMq.InitProducer(*cfg)
	// run http server

	//go func() {
	//	data := GetOrder()
	//	PublishMessage(&data)
	//}()

	//for i := 0; i < 10; i++ {
	//	go func() {
	//		data := GetOrder()
	//		rabbitMq.PublishMessage(&data)
	//	}()
	//}

	go func() {
		for {
			data := GetOrder()
			rabbitMq.PublishMessage(&data)
		}
	}()

	//Run(*cfg)
	RunHttpServer(*cfg)
}
