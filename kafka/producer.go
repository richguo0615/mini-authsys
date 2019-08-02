package kafka

import (
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/richguo0615/mini-authsys/model/db"
	"github.com/richguo0615/mini-authsys/utils/byteParser"
	"log"
)

type ProducerType int

const (
	ProducerTypeNone  ProducerType = 0
	ProducerTypeSync  ProducerType = 1
	ProducerTypeAsync ProducerType = 2
)

const (
	ProducerTrans string = "trans"
)

var (
	syncProducers  = make(map[string]sarama.SyncProducer)
	asyncProducers = make(map[string]sarama.AsyncProducer)
)

func InitProducer(name string, brokers []string, producerType ProducerType) (success bool) {
	switch producerType {
	case ProducerTypeAsync:
		if asyncProducers[name] == nil {
			asyncProducers[name] = newAsyncProducer(producerConfig, brokers)
			fmt.Println("kafka new a async producer:", name,"...")
			return true
		}
	case ProducerTypeSync:
		if syncProducers[name] == nil {
			syncProducers[name] = newSyncProducer(producerConfig, brokers)
			fmt.Println("kafka new a sync producer:", name,"...")
			return true
		}
	}
	return false
}

func newAsyncProducer(config *sarama.Config, address [] string) sarama.AsyncProducer {
	producer, err := sarama.NewAsyncProducer(address, config)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	go func(p sarama.AsyncProducer) {
		for {
			select {
			case suc := <-p.Successes():
				fmt.Println("offset: ", suc.Offset, ", timestamp: ", suc.Timestamp.String(), ", partitions: ", suc.Partition)
			case fail := <-p.Errors():
				fmt.Println("err: ", fail.Err)
			}
		}
	}(producer)

	return producer
}

func newSyncProducer(config *sarama.Config, address []string) sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return producer
}

func SendTransIn(trans db.Transaction) (err error) {

	producer := syncProducers[ProducerTrans]

	data, err := byteParser.Encode(trans)
	if err != nil {
		err = errors.New(fmt.Sprintf("kafka producer - %s, data to byte err: %s \n", TopicTransIn, err))
	}

	msg := &sarama.ProducerMessage{
		Topic:     TopicTransIn,
		Value:     sarama.ByteEncoder(data),
	}

	part, offset, err := producer.SendMessage(msg)
	if err != nil {
		err = errors.New(fmt.Sprintf("kafka producer - %s, send msg err = %s \n", TopicTransIn, err))
		log.Print(err)
		return
	} else {
		log.Printf("kafka producer - %s, send msg success. parition: %d, offset: %d", TopicTransIn, part, offset)
	}
	return
}
