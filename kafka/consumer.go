package kafka

import (
	"fmt"
	"github.com/bsm/sarama-cluster"
	"os"
	"os/signal"
)

const (
	ConsumeGroupTrans = "transGroup"
)

const (
	ConsumerTrans = "trans"
)

var (
	consumers = make(map[string]*cluster.Consumer)
)

func InitConsumer(name string, brokers []string) (success bool) {
	if consumers[name] == nil {
		consumers[name] = NewConsumer(consumerConfig, brokers)
		fmt.Println("kafka new a consumer:", name,"...")
		return true
	}
	return false
}

func NewConsumer(config *cluster.Config, brokers [] string) *cluster.Consumer {
	consumer, err := cluster.NewConsumer(brokers, ConsumeGroupTrans, []string{TopicTransIn}, config)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return consumer
}

func ConsumeTransIn() {
	consumer := consumers[ConsumerTrans]

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		fmt.Println("start consume TransIn ...")
		select {
		case msg, ok := <- consumer.Messages():
			if ok {
				fmt.Println("kafka consumer -", TopicTransIn ,"- msg offset: ", msg.Offset, " partition: ", msg.Partition, " timestrap: ", msg.Timestamp.Format("2006-Jan-02 15:04"), " value: ", string(msg.Value))
			}
		case <-signals:
			fmt.Println("consume signals return")
			return
		}
	}
}
