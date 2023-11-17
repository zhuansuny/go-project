package models

import (
	"github.com/jmoiron/sqlx"
	clientV3 "go.etcd.io/etcd/client/v3"
)

var (
	Db             *sqlx.DB
	EtcdClient     *clientV3.Client
	EtcdPrefix     string
	EtcdProductKey string
)

// 从主函数初始化中获取配置
func Init(db *sqlx.DB, etcdClient *clientV3.Client, prefix, productKey string) (err error) {
	Db = db
	EtcdClient = etcdClient
	EtcdPrefix = prefix
	EtcdProductKey = productKey
	return
}
