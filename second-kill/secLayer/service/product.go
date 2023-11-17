package service

import (
	"context"
	"encoding/json"

	"time"

	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

//导入产品信息从etcd
func loadProductFromEtcd(conf *SecLayerConf) (err error) {

	logs.Debug("start get from etcd succ")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	//从初始化好的etcdClient中读取EtcdSecProductKey的配置
	resp, err := secLayerContext.etcdClient.Get(ctx, conf.EtcdConfig.EtcdSecProductKey)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err:%v", conf.EtcdConfig.EtcdSecProductKey, err)
		return
	}

	logs.Debug("get from etcd succ, resp:%v", resp)
	//将读取的数据反序列化到secProductInfo
	var secProductInfo []SecProductInfoConf
	for k, v := range resp.Kvs {
		logs.Debug("key[%v] valud[%v]", k, v)
		err = json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			logs.Error("Unmarshal sec product info failed, err:%v", err)
			return
		}

		logs.Debug("sec info conf is [%v]", secProductInfo)
	}
	//将反序列化之后的商品信息更新
	updateSecProductInfo(conf, secProductInfo)
	logs.Debug("update product info succ, data:%v", secProductInfo)

	//开启协程监控etcd变化
	initSecProductWatcher(conf)

	logs.Debug("init etcd watcher succ")
	return
}

//更新商品信息
func updateSecProductInfo(conf *SecLayerConf, secProductInfo []SecProductInfoConf) {

	var tmp map[int]*SecProductInfoConf = make(map[int]*SecProductInfoConf, 1024)
	for _, v := range secProductInfo {
		produtInfo := v
		produtInfo.secLimit = &SecLimit{}
		tmp[v.ProductId] = &produtInfo
	}

	secLayerContext.RWSecProductLock.Lock() //加锁
	conf.SecProductInfoMap = tmp
	secLayerContext.RWSecProductLock.Unlock()
}

func initSecProductWatcher(conf *SecLayerConf) {
	go watchSecProductKey(conf)
}

func watchSecProductKey(conf *SecLayerConf) {

	key := conf.EtcdConfig.EtcdSecProductKey
	logs.Debug("begin watch key:%s", key)
	var err error
	for {
		rch := secLayerContext.etcdClient.Watch(context.Background(), key)
		var secProductInfo []SecProductInfoConf
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
				updateSecProductInfo(conf, secProductInfo)
			}
		}
	}
}
