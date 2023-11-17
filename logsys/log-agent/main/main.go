package main

import (
	"fmt"
	"logsys/log-agent/kafka"
	"logsys/log-agent/tailf"

	"github.com/astaxie/beego/logs"
)

func main() {
	filename := "../config/logagent.conf" //配置文件路径
	err := loadConf("ini", filename)      //调用配置文件读取函数
	if err != nil {
		fmt.Printf("load conf failed, err: %v\n", err)
		panic("load conf failed")
		return
	}

	err = initLogger() //初始化日志准备写入
	if err != nil {
		fmt.Printf("load logger failed, err: %v\n", err)
		panic("load logger failed")
		return
	}
	logs.Debug("load conf success ,config:%v\n", appConfig) //在日志文件中写入

	//collectConf, err := initEctd(appConfig.etcdAddr, appConfig.etcdKey)
	//if err != nil {
	//	logs.Error("init etcd failed, err:%v\n", err)
	//	return
	//}
	collec := tailf.CollectConf{
		LogPath:"E:/project/spiot/src/web/webWindows.exe_log/log.log",
		Topic:"test",
	}
	collectConf := []tailf.CollectConf{}
	collectConf = append(collectConf, collec)
	err = tailf.InitTail(collectConf, appConfig.chanSize) //初始化tail查看日志
	if err != nil {
		logs.Error("init tail failed, err:%v\n", err)
		return
	}
	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed, err:%v\n", err)
		return
	}

	logs.Debug("initialize success")
	serverRun()
	logs.Info("program exited")

}
