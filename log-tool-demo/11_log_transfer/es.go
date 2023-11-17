package main

import (
	"fmt"

	elastic "gopkg.in/olivere/elastic.v2"
)

type LogMessage struct {
	App     string
	Topic   string
	Message string
}

var (
	esClient *elastic.Client
)

func initEs() (err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(logConfig.ESAddr)) //"http://127.0.0.1:9200/"
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}
	esClient = client
	fmt.Println("connect es succ")
	return

}

func sendToES(topic string, data []byte) (err error) {
	msg := &LogMessage{}
	msg.Topic = topic
	msg.Message = string(data)

	_, err = esClient.Index().
		Index(topic).
		Type(topic).
		//Id(fmt.Sprintf("%d", i)).
		BodyJson(msg).
		Do()
	if err != nil {
		panic(err)
		return
	}
	return

}
