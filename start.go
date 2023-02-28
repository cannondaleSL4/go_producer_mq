package main

import (
	. "go_producer_mq/config"
	. "go_producer_mq/httpserver"
	rabbitMq "go_producer_mq/rabbitMq"
	"log"
	"time"
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

	producer := &rabbitMq.ProducerStruct{}

	go producer.InitProducer(*cfg)

	go func() {
		var counter int = 0
		for {
			if (counter % 10000) == 0 {
				counter = 0
				time.Sleep(10000 * time.Millisecond)
			}
			counter++
			data := GetOrder()
			producer.PublishMessage(&data)
		}
	}()

	RunHttpServer(*cfg)
}
