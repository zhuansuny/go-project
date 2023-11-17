package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

type LogConfig struct {
	KafkaAddr  string //配置文件中kafka的ip加端口
	ESAddr     string //配置文件中es的ip加端口
	LogPath    string //配置文件中日志存储路径
	LogLevel   string //日志级别
	KafkaTopic string //kafka的topic

	etcdAddr string
	etcdKey  string
}

var (
	logConfig *LogConfig //全局变量
)

//---------初始化配置信息，从配置文件中读取信息，传入到logConfig中-------------
func initConfig(confType string, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename) //调用config包中的NewConfig函数，创建一个config实例
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	logConfig = &LogConfig{}
	logConfig.LogLevel = conf.String("logs::log_level") //将从配置文件中读取的log_level传入全局logConfig中
	if len(logConfig.LogLevel) == 0 {
		logConfig.LogLevel = "debug"
	}

	logConfig.LogPath = conf.String("logs::log_path") //同上
	if len(logConfig.LogPath) == 0 {
		logConfig.LogPath = "./logs"
	}

	logConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(logConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	logConfig.ESAddr = conf.String("es::addr")
	if len(logConfig.ESAddr) == 0 {
		err = fmt.Errorf("invalid es addr")
		return
	}

	logConfig.KafkaTopic = conf.String("kafka::topic")
	if len(logConfig.ESAddr) == 0 {
		err = fmt.Errorf("invalid es addr")
		return
	}

	logConfig.etcdAddr = conf.String("etcd::addr")
	if len(logConfig.etcdAddr) == 0 {
		err = fmt.Errorf("invalid etcdAddr ")
		return
	}

	logConfig.etcdKey = conf.String("etcd::configKey")
	if len(logConfig.etcdKey) == 0 {
		err = fmt.Errorf("invalid etcdKey")
		return
	}
	return
}
