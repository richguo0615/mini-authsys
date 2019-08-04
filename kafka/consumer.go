package kafka

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const (
	ConsumeGroupTrans = "transGroup"
)

const (
	ConsumerTrans = "trans"
)

var (
	consumers = make(map[string]*kafka.Consumer)
)

func InitConsumer(name string, brokers []string) (success bool) {
	if consumers[name] == nil {
		consumers[name] = NewConsumer()
		fmt.Println("kafka new a consumer:", name,"...")
		return true
	}
	return false
}

func NewConsumer() *kafka.Consumer {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          ConsumeGroupTrans,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	consumer.SubscribeTopics([]string{TopicTransIn}, nil)

	return consumer
}

func ConsumeTransIn() {
	consumer := consumers[ConsumerTrans]

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
