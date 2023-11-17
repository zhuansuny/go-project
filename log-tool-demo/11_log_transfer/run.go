package main

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

func run() (err error) {
	partitionList, err := kafkaClient.client.Partitions(kafkaClient.topic)
	if err != nil {
		logs.Error("Failed to get the list of partitions: ", err)
		return
	}

	for partition := range partitionList {
		pc, errRed := kafkaClient.client.ConsumePartition(logConfig.KafkaTopic, int32(partition), sarama.OffsetNewest)
		if errRed != nil {
			err = errRed
			logs.Error("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		kafkaClient.wg.Add(1)
		go func(pc sarama.PartitionConsumer) {

			fmt.Println("pc.Messages,")
			for msg := range pc.Messages() {
				logs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				err = sendToES(kafkaClient.topic, msg.Value)
				if err != nil {
					logs.Warn("send to es failed, err:%v", err)
				}

			}
			kafkaClient.wg.Done()
		}(pc)

	}

	kafkaClient.wg.Wait()
	return

}
