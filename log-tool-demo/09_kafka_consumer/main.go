package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	wg sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer(strings.Split("127.0.0.1:9092", ","), nil)
	if err != nil {
		fmt.Println("failed to start consumer: ", err)
		return
	}

	partitionList, err := consumer.Partitions("nginx_log")
	if err != nil {
		fmt.Println("failed to get the list of partitions:", err)
		return
	}
	fmt.Println(partitionList)

	for _, partition := range partitionList {
		fmt.Println(partition)
		pc, err := consumer.ConsumePartition("nginx_log", partition, sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {

			fmt.Println("pc.Messages,")
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			wg.Done()
		}(pc)
	}

	wg.Wait()
	consumer.Close()

}
