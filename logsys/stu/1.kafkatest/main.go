package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)
var (
	client sarama.SyncProducer
)
func main(){
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	cli, err := sarama.NewSyncProducer([]string{"111.229.157.64:9092"}, config)
	client = cli
	if err != nil {
		logs.Error(" init kafka producer failed, err:", err)
		return
	}
	err = SendToKafka("hello","test")
	if err != nil {
		logs.Error(" init kafka producer failed, err:", err)
		cli.Close()
		return
	}
	logs.Info(" send kafka success")
	cli.Close()
}

func SendToKafka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	_, _, err = client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed err:%v data:%v topic:%v\n", err, data, topic)
		return
	}
	//logs.Debug("send succ , pid:%v offset:%v topic:%V\n", pid, offset, topic)

	return

}
