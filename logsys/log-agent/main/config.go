package main //读取配置文件内容

import (
	"errors"
	"fmt"

	"logsys/log-agent/tailf"

	"github.com/astaxie/beego/config"
)

var (
	appConfig *Config //声明一个全局config结构体
)

type Config struct { //定义配置结构体
	//写入日志的配置
	logLevel  string //日志级别
	logPath   string //日志存储路径
	chanSize  int    //读取日志的管道容量
	kafkaAddr string

	collectConf []tailf.CollectConf //读日志collectConf的配置结构体，已经弃用，使用etcd读取配置

	etcdAddr string
	etcdKey  string
}

// type CollectConf struct {
// 	LogPath string
// 	Topic   string
// }

//-----------配置文件读取--之读日志collectConf配置-----
func loadCollectConf(conf config.Configer) (err error) {

	var cc tailf.CollectConf                      //声明一个读日志结构体
	cc.LogPath = conf.String("collect::log_path") //读取配置文件中的日志地址
	if len(cc.LogPath) == 0 {                     //若未读到则报错
		err = errors.New("invalid collect::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic") //读取配置文件中的topic
	if len(cc.LogPath) == 0 {                //若未读到则报错
		err = errors.New("invalid collect::topic")
		return
	}

	appConfig.collectConf = append(appConfig.collectConf, cc) //将collectConf结构体写入到全局变量中
	return

}

//------------配置文件读取--------------
func loadConf(confType, filename string) (err error) {

	conf, err := config.NewConfig(confType, filename) //配置文件的路径
	if err != nil {
		fmt.Println("new config failed err:", err)
		return
	}

	appConfig = &Config{}                               //全局结构体赋值
	appConfig.logLevel = conf.String("logs::log_level") //读取配置文件中的logLevel赋值给全局结构体
	if len(appConfig.logLevel) == 0 {                   //未读取到则使用默认值debug
		appConfig.logLevel = "debug"
	}

	appConfig.logPath = conf.String("logs::log_path") //读取配置文件中的logPath赋值给全局结构体
	if len(appConfig.logPath) == 0 {                  //未读取到则使用默认值
		appConfig.logLevel = "../logs/logagent.log"
	}

	appConfig.kafkaAddr = conf.String("kafka::server_ip")
	if len(appConfig.logPath) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}
	appConfig.etcdAddr = conf.String("etcd::addr")
	if len(appConfig.etcdAddr) == 0 {
		err = fmt.Errorf("invalid etcdAddr ")
		return
	}

	appConfig.etcdKey = conf.String("etcd::configKey")
	if len(appConfig.etcdKey) == 0 {
		err = fmt.Errorf("invalid etcdKey")
		return
	}

	appConfig.chanSize, err = conf.Int("collect::chan_size") //读取chanSize
	if err != nil {
		appConfig.chanSize = 100
	}

	err = loadCollectConf(conf) //调用loadCollectConf函数，读日志collectConf配置
	if err != nil {
		fmt.Printf("load collect conf failed, err:%v\n", err)
		return
	}
	return

}
