package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
)

var (
	producerConfig *sarama.Config
	consumerConfig *cluster.Config
)

const (
	TopicTransIn string = "transIn"
)

func InitConfig() {
	InitProducerConfig()
	InitConsumerConfig()
}

func InitProducerConfig() {
	fmt.Println("kafka init producer config...")
	producerConfig = sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Return.Successes = true
	producerConfig.Producer.Return.Errors = true
	producerConfig.Version = sarama.V0_11_0_0
}

func InitConsumerConfig() {
	fmt.Println("kafka init consumer config...")
	consumerConfig = cluster.NewConfig()
	consumerConfig.Group.Mode = cluster.ConsumerModePartitions
	consumerConfig.Group.Return.Notifications = true
	//consumerConfig.Group.PartitionStrategy = cluster.StrategyRoundRobin
}
