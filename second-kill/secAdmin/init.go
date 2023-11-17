package main

import (
	"encoding/json"
	"fmt"
	"secKill/secAdmin/models"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/astaxie/beego/logs"
	clientV3 "go.etcd.io/etcd/client/v3"
)

var (
	Db         *sqlx.DB
	EtcdClient *clientV3.Client
)

// 初始化mysql数据库
func initDb() (err error) {
	database, err := sqlx.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/seckill?parseTime=true")
	if err != nil {
		logs.Warn("open mysql failed,", err)
		return
	}
	Db = database
	logs.Debug("connect to mysql succ")
	return
}

// -------------------初始化etcd--------------------
func initEtcd() (err error) {
	cli, err := clientV3.New(clientV3.Config{
		Endpoints:   []string{AppConf.etcdConf.Addr},
		DialTimeout: time.Duration(AppConf.etcdConf.Timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}

	EtcdClient = cli
	logs.Debug("connect to etcd succ")
	return
}

// -------------------日志级别转换--------------------
func convertLogLevel(level string) int {

	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}

	return logs.LevelDebug
}

// -------------------初始化日志--------------------
func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = AppConf.logsConf.LogPath
	config["level"] = convertLogLevel(AppConf.logsConf.LogLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

// -------------------进行总的初始化工作--------------------
func initAll() (err error) {
	err = initConfig()
	if err != nil {
		logs.Warn("init config failed, err:%v", err)
		return
	}
	err = initLogger()
	if err != nil {
		logs.Warn("init logger failed, err:%v", err)
		return
	}
	err = initDb()
	if err != nil {
		logs.Warn("init Db failed, err:%v", err)
		return
	}

	err = initEtcd()
	if err != nil {
		logs.Warn("init etcd failed, err:%v", err)
		return
	}

	err = models.Init(Db, EtcdClient, AppConf.etcdConf.EtcdKeyPrefix, AppConf.etcdConf.ProductKey)
	if err != nil {
		logs.Warn("init model failed, err:%v", err)
		return
	}
	return
}
