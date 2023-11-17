package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./logs/logcollect.log" //创建日志文件
	config["level"] = logs.LevelDebug            //日志级别LevelDebug可以全部打印（7），LevelInfo(6)打印Warn

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))

	logs.Debug("this is a test, my name is %s", "stu01")
	logs.Trace("this is a trace, my name is %s", "stu02")
	logs.Warn("this is a warn, my name is %s", "stu03")

}
