package main //创建日志文件

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

//将从配置文件中读取的log_level转换成可识别的INT
func convertLogLevel(level string) int {
	switch level {
	case "dubug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug //默认为debug  (7)
}

//----
func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = logConfig.LogPath //从全局结构体中读取路径
	config["level"] = convertLogLevel(logConfig.LogLevel)

	configStr, err := json.Marshal(config) //将config map 序列化
	if err != nil {
		fmt.Println(" initlogger failed,marshal err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr)) //日志开始准备写入

	return

}
