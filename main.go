package main

import (
	"github.com/richguo0615/mini-authsys/conf"
	"github.com/richguo0615/mini-authsys/kafka"
	"github.com/richguo0615/mini-authsys/router"
)

func main() {
	conf.InitDB()
	conf.InitRedis()
	kafka.InitConfig()
	kafka.InitProducer(kafka.ProducerTrans, []string{"localhost:9092"}, kafka.ProducerTypeSync)
	kafka.InitConsumer(kafka.ConsumerTrans, []string{"localhost:9092"})

	go kafka.ConsumeTransIn()

	router.InitRoute()
}
