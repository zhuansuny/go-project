package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var wait sync.WaitGroup

//---------------程序主任务，开启协程读取kafka分区中的数据，并发送给es----------
func run() (err error) {

	for _, conf := range collectConf {
		wait.Add(1)
		go createNewTask(conf) //每一个程序主任务启动
	}

	wait.Wait()
	return
}

func createNewTask(conf *CollectConf) {
	conf.exitChan = make(chan int, 1) //必须要先分配空间
	fmt.Println("开启了一个新的协程读取", conf.LogPath, conf)
	logs.Debug("开启了一个新的协程读取", conf.LogPath)
	kafkaClient.topic = conf.Topic
	partitionList, err := kafkaClient.client.Partitions(kafkaClient.topic) //读取kafka topic日志的分区
	if err != nil {
		logs.Error("Failed to get the list of partitions: ", err)
		return
	}

	for partition := range partitionList { //每个分区开启一个协程
		pc, errRed := kafkaClient.client.ConsumePartition(conf.Topic, int32(partition), sarama.OffsetNewest)
		if errRed != nil {
			err = errRed
			logs.Error("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		kafkaClient.wg.Add(1) //sync.WaitGroup加1
		go sendAndExit(pc, conf)

	}

	kafkaClient.wg.Wait() //主线程一直阻塞在这，直到sync.WaitGroup为0
	wait.Done()
	return
}

func sendAndExit(pc sarama.PartitionConsumer, conf *CollectConf) {
	for {
		select {
		case msg := <-pc.Messages():
			logs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			err := sendToES(kafkaClient.topic, msg.Value)
			if err != nil {
				logs.Warn("send to es failed, err:%v", err)
			}

		case <-conf.exitChan:
			logs.Warn("tail obj will exited, conf:%v", conf.LogPath)
			fmt.Printf("协程%v退出\n", conf.LogPath)
			kafkaClient.wg.Done() //sync.WaitGroup减1
			return
		}

	}

}
