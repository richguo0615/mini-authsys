package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	producerConfig *sarama.Config
)

const (
	TopicTransIn string = "transIn"
)

func InitConfig() {
	InitProducerConfig()
}

func InitProducerConfig() {
	fmt.Println("kafka init producer config...")
	producerConfig = sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Return.Successes = true
	producerConfig.Producer.Return.Errors = true
	producerConfig.Version = sarama.V0_11_0_0
}