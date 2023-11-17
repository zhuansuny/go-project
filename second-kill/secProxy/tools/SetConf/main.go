package main

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

const (
	EtcdKey = "/seckill/secProxy/product"
)

type SecInfoConf struct {
	ProductId         int
	StartTime         int
	EndTime           int
	Status            int
	Total             int
	Left              int
	OnePersonBuyLimit int
	BuyRate           float64
	//每秒最多能卖多少个
	SoldMaxLimit int
}

func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	var SecInfoConfArr []SecInfoConf
	SecInfoConfArr = append(
		SecInfoConfArr,
		SecInfoConf{
			ProductId:         1029,
			StartTime:         1569899504,
			EndTime:           1569903104,
			Status:            0,
			Total:             1000,
			Left:              1000,
			OnePersonBuyLimit: 2,
			BuyRate:           0.1,
			SoldMaxLimit:      10,
		},
	)
	SecInfoConfArr = append(
		SecInfoConfArr,
		SecInfoConf{
			ProductId:         1027,
			StartTime:         1569899504,
			EndTime:           1569903104,
			Status:            0,
			Total:             2000,
			Left:              1000,
			OnePersonBuyLimit: 2,
			BuyRate:           0.1,
			SoldMaxLimit:      10,
		},
	)

	data, err := json.Marshal(SecInfoConfArr)
	if err != nil {
		fmt.Println("json failed, ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cli.Delete(ctx, EtcdKey)
	//return
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func main() {
	SetLogConfToEtcd()
}
