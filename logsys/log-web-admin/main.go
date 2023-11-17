package main

import (
	"fmt"
	_ "logsys/log-web-admin/routers"
	"time"

	"github.com/astaxie/beego/logs"

	_ "github.com/go-sql-driver/mysql"

	"logsys/log-web-admin/models"

	"github.com/astaxie/beego"
	"github.com/jmoiron/sqlx"
	etcdclient "go.etcd.io/etcd/clientv3"
)

//初始数据库
func initDb() (err error) {
	database, err := sqlx.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/logadmin?parseTime=true")
	if err != nil {
		logs.Warn("open mysql failed,", err)
		return
	}

	models.InitDb(database)
	return
}

//初始化Etcd
func initEtcd() (err error) {
	cli, err := etcdclient.New(etcdclient.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	models.InitEtcd(cli)
	return
}

func main() {
	err := initDb() //初始数据库
	if err != nil {
		logs.Warn("initDb failed, err:%v", err)
		return
	}

	err = initEtcd() //初始化Etcd
	if err != nil {
		logs.Warn("initEtcd failed, err:%v", err)
	}
	beego.Run() //web服务以及监听在9091
}
