package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"go.etcd.io/etcd/mvcc/mvccpb"

	"github.com/astaxie/beego/logs"
	etcd_client "go.etcd.io/etcd/clientv3"
)

type CollectConf struct {
	LogPath  string `json:"logPath"`
	Topic    string `json:"topic"`
	exitChan chan int
}

// type CollectConfMgr struct {
// 	collectConf []*CollectConf
// 	exitChan    chan int
// }

type EtcdClient struct {
	client *etcd_client.Client
	keys   []string
}

var (
	etcdClient  *EtcdClient
	collectConf []*CollectConf
)

func initEtcd() (err error) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{logConfig.etcdAddr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}
	logs.Info("connect etcd succ")

	etcdClient = &EtcdClient{
		client: cli,
	}

	if strings.HasSuffix(logConfig.etcdKey, "/") == false { //若读取的配置不是以/结尾就加上/
		logConfig.etcdKey = logConfig.etcdKey + "/"
	}

	//for _, ip := range localIPArray {
	etcdKey := fmt.Sprintf("%s%s", logConfig.etcdKey, "127.0.0.1") //将传入的key加上IP地址，当做etcd的Key
	etcdClient.keys = append(etcdClient.keys, etcdKey)

	logs.Debug("Etcd load succ etcdKey:%s", etcdKey)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, etcdKey) //超过一秒就取消, 将读取的kv数据放入到resp中
	if err != nil {
		logs.Error("client get from etcd failed, err:%v", err)
		return
	}
	cancel() //成功后释放context资源
	logs.Debug("resp from etcd:%v", resp.Kvs)

	for _, v := range resp.Kvs { //从resp中取出KV
		if string(v.Key) == etcdKey { //判断key是否等于etcdKey
			err = json.Unmarshal(v.Value, &collectConf) //将value反序列化,放入到全局collectConf中
			if err != nil {
				logs.Error("unmarshal failed, err:%v", err)
				continue
			}

			logs.Debug("log config is %v", collectConf)

		}
	}
	//}
	initEtcdWatcher()
	return
}

func initEtcdWatcher() { //初始化Etcd 监管 key

	for _, key := range etcdClient.keys {
		go watchKey(key)
	}
}

func watchKey(key string) { //通过更改etcd的值，直接更改路径文件，不需要重启服务
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{logConfig.etcdAddr},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect etcd failed, err:", err)
		return
	}

	logs.Debug("begin watch key :%s", key)

	for {
		rch := cli.Watch(context.Background(), key) //检测节点的变化
		var getConfSucc = true
		var colConf []*CollectConf
		for wresp := range rch { //输出变化
			for _, v := range wresp.Events {
				if v.Type == mvccpb.DELETE {
					logs.Warn("key[%s] Config delete", key)
					continue
				}
				if v.Type == mvccpb.PUT && string(v.Kv.Key) == key {
					err = json.Unmarshal(v.Kv.Value, &colConf)
					if err != nil {
						logs.Error("key [%s], Unmarshal[%s], err:%v ", err)
						getConfSucc = false
						continue
					}

				}
				logs.Debug("get config from etcd, %s %q : %q\n", v.Type, v.Kv.Key, v.Kv.Value)
			}
			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", colConf)

				UpdateConf(colConf)

			}
		}

	}

}

func UpdateConf(colConf []*CollectConf) {
	for _, oneConf := range colConf {
		var isRunning = false
		for _, conf := range collectConf {
			if oneConf.LogPath == conf.LogPath {
				isRunning = true
				break
			}
		}

		if isRunning {
			continue
		}
		wait.Add(1)
		go createNewTask(oneConf)
		collectConf = append(collectConf, oneConf)
	}

	var col []*CollectConf
	for _, conf := range collectConf {
		status := 0
		for _, oneConf := range colConf {
			if oneConf.LogPath == conf.LogPath {
				status = 1
				break
			}
		}

		if status == 0 {
			conf.exitChan <- 1
			fmt.Println("exitChan加1", conf)
			continue
		}
		col = append(col, conf)
	}

	collectConf = col

}
