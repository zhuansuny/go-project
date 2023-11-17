package main

import (
	"github.com/astaxie/beego/logs"
)

func main() {
	err := initConfig("ini", "./conf/log_transfer.conf") //初始化配置函数，传入配置文件类型，以及路径
	if err != nil {
		panic(err)
		return
	}

	err = initEtcd()
	if err != nil {
		panic(err)
		return
	}

	err = initLogger() //初始化日志相关
	if err != nil {
		panic(err)
		return
	}
	logs.Debug("init logger succ")

	err = initKafka() //初始化Kafka
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}
	logs.Debug("init kafka succ")

	err = initEs() //初始化elastic search
	if err != nil {
		logs.Error("init es failed, err:%v", err)
		return
	}

	logs.Debug("init es client succ")

	err = run() //程序主任务启动
	if err != nil {
		logs.Error("run  failed, err:%v", err)
		return
	}

	logs.Warn("warning, log_transfer is exited")

}
