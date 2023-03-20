package main

import (
	"fmt"
	. "go_producer_mq/config"
	. "go_producer_mq/httpserver"
	rabbitMq "go_producer_mq/rabbitMq"
	"log"
	"os"
)

func main() {

	cfgPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	producer := &rabbitMq.ProducerStruct{}

	ok, _ := producer.QueueExists(*cfg)

	if !ok {
		fmt.Println("queue does not exist. program execution stops.")
		os.Exit(1)
	}

	go producer.InitProducer(*cfg)

	go func() {
		//var counter int = 0
		for {
			//if (counter % 10000) == 0 {
			//	counter = 0
			//	time.Sleep(1000 * time.Millisecond)
			//}
			//counter++
			data := GetOrder()
			producer.PublishMessage(&data)
			//time.Sleep(10000 * time.Millisecond)
			//time.Sleep(1000 * time.Millisecond)
		}
	}()

	//go func() {
	//	//var counter int = 0
	//	for {
	//		//if (counter % 10000) == 0 {
	//		//	counter = 0
	//		//	time.Sleep(1000 * time.Millisecond)
	//		//}
	//		//counter++
	//		data := GetOrder()
	//		producer.PublishMessage(&data)
	//		//time.Sleep(1000 * time.Millisecond)
	//	}
	//}()
	//
	//go func() {
	//	//var counter int = 0
	//	for {
	//		//if (counter % 10000) == 0 {
	//		//	counter = 0
	//		//	time.Sleep(1000 * time.Millisecond)
	//		//}
	//		//counter++
	//		data := GetOrder()
	//		producer.PublishMessage(&data)
	//		//time.Sleep(1000 * time.Millisecond)
	//	}
	//}()
	//
	//go func() {
	//	//var counter int = 0
	//	for {
	//		//if (counter % 10000) == 0 {
	//		//	counter = 0
	//		//	time.Sleep(1000 * time.Millisecond)
	//		//}
	//		//counter++
	//		data := GetOrder()
	//		producer.PublishMessage(&data)
	//		//time.Sleep(1000 * time.Millisecond)
	//	}
	//}()
	//
	//go func() {
	//	//var counter int = 0
	//	for {
	//		//if (counter % 10000) == 0 {
	//		//	counter = 0
	//		//	time.Sleep(1000 * time.Millisecond)
	//		//}
	//		//counter++
	//		data := GetOrder()
	//		producer.PublishMessage(&data)
	//		//time.Sleep(1000 * time.Millisecond)
	//	}
	//}()
	//
	//go func() {
	//	//var counter int = 0
	//	for {
	//		//if (counter % 10000) == 0 {
	//		//	counter = 0
	//		//	time.Sleep(1000 * time.Millisecond)
	//		//}
	//		//counter++
	//		data := GetOrder()
	//		producer.PublishMessage(&data)
	//		//time.Sleep(1000 * time.Millisecond)
	//	}
	//}()

	RunHttpServer(*cfg)
}
