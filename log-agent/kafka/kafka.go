package kafka //发送日志给kafka

import (
	"github.com/astaxie/beego/logs"

	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error(" init kafka producer failed, err:", err)
		return
	}

	logs.Debug("init kafka success")
	return
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
