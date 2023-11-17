package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

func main() {

	conf, err := config.NewConfig("ini", "./logagent.conf") //配置文件的路径
	if err != nil {
		fmt.Println("new config failed err:", err)
		return
	}
	log_level, err := conf.Int("logs::log_level")
	if err != nil {
		fmt.Println("read config failed,err :", err)
		return
	}

	config1 := make(map[string]interface{})
	config1["filename"] = "./logs/logcollect.log" //创建日志文件
	config1["level"] = log_level                  //7

	configStr, err := json.Marshal(config1)
	if err != nil {
		fmt.Println("marshal failed", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))

	logs.Debug("this is a test, my name is %s", "stu01")
	logs.Trace("this is a trace, my name is %s", "stu02")
	logs.Warn("this is a warn, my name is %s", "stu03")

}
