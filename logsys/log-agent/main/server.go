package main

import (
	"fmt"
	"logsys/log-agent/kafka"
	"logsys/log-agent/tailf"
	"time"

	"github.com/astaxie/beego/logs"
)

func serverRun() {
	for {
		msg := tailf.GetOneLine()
		err := sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed,err:%v\n", err)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("send kafka sucess", msg.Topic, msg.Msg)
		//time.Sleep(time.Second) //休眠一秒，让系统一秒发送一行
	}
	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	logs.Info("read msg:%s, topic:%s\n", msg.Msg, msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}
