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
	kafka.InitProducer(kafka.ProducerTrans, []string{"172.16.132.16:9092"}, kafka.ProducerTypeSync)
	kafka.InitConsumer(kafka.ConsumerTrans, []string{"172.16.132.16:9092"})

	go kafka.ConsumeTransIn()

	router.InitRoute()
}
