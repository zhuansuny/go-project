package main

import (
	"context"
	"fmt"
	"time"

	etcd_client "go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	cli.Put(context.Background(), "/logagent/conf/", "2314")  //写入
	rch := cli.Watch(context.Background(), "/logagent/conf/") //检测节点的变化
	for wresp := range rch {                                  //输出变化
		for _, v := range wresp.Events {
			fmt.Printf("%s %q :%q \n", v.Type, v.Kv.Key, v.Kv.Value)
		}
	}

}
