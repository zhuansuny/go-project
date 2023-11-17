package main

import (
	"strings"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

type KafkaClient struct {
	client sarama.Consumer
	addr   string
	topic  string
	wg     sync.WaitGroup
}

var (
	kafkaClient *KafkaClient
)

func initKafka() (err error) {
	kafkaClient = &KafkaClient{}

	consumer, err := sarama.NewConsumer(strings.Split(logConfig.KafkaAddr, ","), nil)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	kafkaClient.client = consumer
	kafkaClient.addr = logConfig.KafkaAddr
	kafkaClient.topic = logConfig.KafkaTopic
	return

}
