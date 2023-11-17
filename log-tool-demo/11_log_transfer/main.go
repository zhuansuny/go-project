package main

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	err := initConfig("ini", "./conf/log_transfer.conf")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(logConfig)

	err = initLogger()
	if err != nil {
		panic(err)
		return
	}
	logs.Debug("init logger succ")

	err = initKafka()
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}
	logs.Debug("init kafka succ")

	err = initEs()
	if err != nil {
		logs.Error("init es failed, err:%v", err)
		return
	}

	logs.Debug("init es client succ")

	err = run()
	if err != nil {
		logs.Error("run  failed, err:%v", err)
		return
	}

	logs.Warn("warning, log_transfer is exited")

}
