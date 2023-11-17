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

//---------------初始化kafka,设置kafka地址、topic------------------
func initKafka() (err error) {
	kafkaClient = &KafkaClient{}

	consumer, err := sarama.NewConsumer(strings.Split(logConfig.KafkaAddr, ","), nil)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	kafkaClient.client = consumer
	kafkaClient.addr = logConfig.KafkaAddr

	//由于没有直接连接etcd无法预先读取topic，需要手动更改配置文件的topic
	//kafkaClient.topic = logConfig.KafkaTopic //Topic固定为nginx_log，若要更改topic,从配置文件中更改
	//kafkaClient.topic = logConfig.KafkaTopic //topic更改为从etcd中读取
	return

}
