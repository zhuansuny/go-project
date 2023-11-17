package service

import (
	"time"

	"github.com/astaxie/beego/logs"
	clientV3 "go.etcd.io/etcd/client/v3"
)

// 总初始化
func InitSecLayer(conf *SecLayerConf) (err error) {
	//初始化redis
	err = initRedis(conf)
	if err != nil {
		logs.Error("init redis failed, err:%v", err)
		return
	}

	logs.Debug("init redis succ")
	//初始化etcd
	err = initEtcd(conf)
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}

	logs.Debug("init etcd succ")
	//导入etcd配置，并监听
	err = loadProductFromEtcd(conf)
	if err != nil {
		logs.Error("load product from etcd failed, err:%v", err)
		return
	}

	logs.Debug("load product succ")
	secLayerContext.secLayerConf = conf
	secLayerContext.Read2HandleChan = make(chan *SecRequest,
		secLayerContext.secLayerConf.Read2handleChanSize)

	secLayerContext.Handle2WriteChan = make(chan *SecResponse,
		secLayerContext.secLayerConf.Handle2WriteChanSize)

	secLayerContext.HistoryMap = make(map[int]*UserBuyHistory, 1000000)

	secLayerContext.productCountMgr = NewProductCountMgr()

	logs.Debug("init all succ")
	return
}

// 初始化etcd,并创建实例放入secLayerContext中
func initEtcd(conf *SecLayerConf) (err error) {
	cli, err := clientV3.New(clientV3.Config{
		Endpoints:   []string{conf.EtcdConfig.EtcdAddr},
		DialTimeout: time.Duration(conf.EtcdConfig.Timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}

	secLayerContext.etcdClient = cli
	logs.Debug("init etcd succ")
	return
}
