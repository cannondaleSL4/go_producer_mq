package main

import (
	"fmt"
	. "go_producer_mq/config"
	. "go_producer_mq/httpserver"
	. "go_producer_mq/rabbitMq"
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

	//InitProducer(*cfg)

	// Run the http server

	InitProducer(*cfg)

	Run(*cfg)

	fmt.Println("test")
}
