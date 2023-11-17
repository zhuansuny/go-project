package main

import (
	"context"
	"encoding/json"
	"fmt"
	"secKill/secProxy/service"
	"time"

	"go.etcd.io/etcd/mvcc/mvccpb"

	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	clientV3 "go.etcd.io/etcd/client/v3"
)

var (
	redisPool  *redis.Pool
	etcdClient *clientV3.Client
)

//-------------------初始化redis--------------------
// func initRedis() (err error) {
// 	redisPool = &redis.Pool{
// 		MaxIdle:     secKillConf.RedisConf.RedisMaxIdle,
// 		MaxActive:   secKillConf.RedisConf.RedisMaxActive,
// 		IdleTimeout: time.Duration(secKillConf.RedisConf.RedisIdleTimeout) * time.Second,
// 		Dial: func() (redis.Conn, error) {
// 			return redis.Dial("tcp", secKillConf.RedisConf.RedisAddr)
// 		},
// 	}

// 	conn := redisPool.Get()
// 	defer conn.Close()

// 	_, err = conn.Do("ping")
// 	if err != nil {
// 		logs.Error("ping redis failed, err:%v", err)
// 		return
// 	}

// 	return
// }

// -------------------初始化etcd--------------------
func initEtcd() (err error) {
	cli, err := clientV3.New(clientV3.Config{
		Endpoints:   []string{secKillConf.EtcdConf.EtcdAddr},
		DialTimeout: time.Duration(secKillConf.EtcdConf.Timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}

	etcdClient = cli
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
	config["filename"] = secKillConf.LogPath
	config["level"] = convertLogLevel(secKillConf.LogLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

// -------------------程序运行时从etcd中读取配置--------------------
func loadSecConf() (err error) {

	resp, err := etcdClient.Get(context.Background(), secKillConf.EtcdConf.EtcdSecProductKey)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err:%v", secKillConf.EtcdConf.EtcdSecProductKey, err)
		return
	}

	var secProductInfo []service.SecProductInfoConf
	for k, v := range resp.Kvs {
		logs.Debug("key[%v] valud[%v]", k, v)
		err = json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			logs.Error("Unmarshal sec product info failed, err:%v", err)
			return
		}

		logs.Debug("sec info conf is [%v]", secProductInfo)
	}

	updateSecProductInfo(secProductInfo)
	return
}

// -------------------进行总的初始化工作--------------------
func initSec() (err error) {

	err = initLogger()
	if err != nil {
		logs.Error("init logger failed, err:%v", err)
		return
	}
	/*
		err = initRedis()
		if err != nil {
			logs.Error("init redis failed, err:%v", err)
			return
		}
	*/
	err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}

	err = loadSecConf()
	if err != nil {
		logs.Error("load sec conf failed, err:%v", err)
		return
	}

	service.InitService(secKillConf)
	initSecProductWatcher()

	logs.Info("init sec succ")
	return
}

// -------------------etcd监控--------------------
func initSecProductWatcher() {
	go watchSecProductKey(secKillConf.EtcdConf.EtcdSecProductKey)
}

// -------------------etcd监控--------------------
func watchSecProductKey(key string) {

	cli, err := clientV3.New(clientV3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}

	logs.Debug("begin watch key:%s", key)
	for {
		rch := cli.Watch(context.Background(), key)
		var secProductInfo []service.SecProductInfoConf
		var getConfSucc = true

		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config deleted", key)
					continue
				}

				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &secProductInfo)
					if err != nil {
						logs.Error("key [%s], Unmarshal[%s], err:%v ", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd, %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}

			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", secProductInfo)
				updateSecProductInfo(secProductInfo)
			}
		}

	}
}

// -----------------------更新product配置信息-------------
func updateSecProductInfo(secProductInfo []service.SecProductInfoConf) {
	//先将配置读取到tmp中间变量在加锁写入
	var tmp map[int]*service.SecProductInfoConf = make(map[int]*service.SecProductInfoConf, 100)
	for _, v := range secProductInfo {
		productInfo := v
		tmp[v.ProductId] = &productInfo
	}

	secKillConf.RWSecProductLock.Lock() //加锁
	secKillConf.SecProductInfoMap = tmp
	secKillConf.RWSecProductLock.Unlock()
	fmt.Println("配置读取为", secKillConf.SecProductInfoMap)
}
