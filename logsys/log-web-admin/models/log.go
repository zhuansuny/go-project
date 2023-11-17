package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/astaxie/beego/logs"
	etcdclient "go.etcd.io/etcd/clientv3"
)

var (
	etcdClient *etcdclient.Client
)

var logConfArr []CollectConf

type LogInfo struct {
	AppId      int    `db:"app_id"`
	AppName    string `db:"app_name"`
	LogId      int    `db:"log_id"`
	Status     int    `db:"status"`
	CreateTime string `db:"create_time"`
	LogPath    string `db:"log_path"`
	Topic      string `db:"topic"`
}

func InitEtcd(client *etcdclient.Client) {
	etcdClient = client
}

type CollectConf struct {
	LogPath string `json:"logpath"`
	Topic   string `json:"topic"`
}

//---------获得数据库中所有日志数据---------
func GetAllLogInfo() (loglist []LogInfo, err error) {
	err = Db.Select(&loglist,
		"select a.app_id, b.app_name, a.create_time, a.topic, a.log_id, a.status, a.log_path from tbl_log_info a, tbl_app_info b where a.app_id=b.app_id")
	if err != nil {
		logs.Warn("Get All App Info failed, err:%v", err)
		return
	}

	return
}

//更新数据数据库数据
func UpdateLog(info *LogInfo) (err error) {

	conn, err := Db.Begin()
	if err != nil {
		logs.Warn("UpdateLog failed, Db.Begin error:%v", err)
		return
	}

	defer func() {
		if err != nil {
			conn.Rollback()
			return
		}

		conn.Commit()
	}()

	var appId []int
	err = Db.Select(&appId, "select app_id from tbl_app_info where app_name=?", info.AppName)
	if err != nil || len(appId) == 0 {
		logs.Warn("select app_id failed, Db.Exec error:%v", err)
		return
	}
	info.AppId = appId[0]
	r, err := conn.Exec("update tbl_log_info set log_path=?, topic=?,create_time=?, app_name=? where log_id =? ",
		info.LogPath, info.Topic, info.CreateTime, info.AppName, info.LogId)

	if err != nil {
		logs.Warn("UpdateLog failed, Db.Exec error:%v", err)
		return
	}

	_, err = r.LastInsertId()
	if err != nil {
		logs.Warn("UpdateLog failed, Db.LastInsertId error:%v", err)
		return
	}

	return
}

//创建log数据库数据
func CreateLog(info *LogInfo) (err error) {

	conn, err := Db.Begin()
	if err != nil {
		logs.Warn("CreateLog failed, Db.Begin error:%v", err)
		return
	}

	defer func() {
		if err != nil {
			conn.Rollback()
			return
		}

		conn.Commit()
	}()

	var appId []int
	err = Db.Select(&appId, "select app_id from tbl_app_info where app_name=?", info.AppName)
	if err != nil || len(appId) == 0 {
		logs.Warn("select app_id failed, Db.Exec error:%v", err)
		return
	}

	info.AppId = appId[0]
	r, err := conn.Exec("insert into tbl_log_info(app_id, log_path, topic,create_time, app_name)values(?, ?, ?,?,?)",
		info.AppId, info.LogPath, info.Topic, info.CreateTime, info.AppName)

	if err != nil {
		logs.Warn("CreateApp failed, Db.Exec error:%v", err)
		return
	}

	_, err = r.LastInsertId()
	if err != nil {
		logs.Warn("CreateApp failed, Db.LastInsertId error:%v", err)
		return
	}

	return
}

//删除数据数据库数据
func DeleteLog(info *LogInfo) (err error) {

	r, err := Db.Exec("delete from tbl_log_info where log_id=? ", info.LogId)

	if err != nil {
		logs.Warn("delete failed, Db.Exec error:%v", err)
		return
	}

	_, err = r.LastInsertId()
	if err != nil {
		logs.Warn("delete failed, Db.LastInsertId error:%v", err)
		return
	}

	return
}

//-------将数据库配置发送到etcd中-------------
func SetLogConfToEtcd(etcdKey string) (err error) {

	GetAllLogCof() // 获取数据库配置，并传入全局变量logConfArr，准备传给Etcd

	// logConfArr = append( //将新增加的数据添加上
	// 	logConfArr,
	// 	CollectConf{
	// 		LogPath: info.LogPath,
	// 		Topic:   info.Topic,
	// 	},
	// )

	data, err := json.Marshal(logConfArr)
	if err != nil {
		logs.Warn("marshal failed, err:%v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//cli.Delete(ctx, EtcdKey)
	//return
	_, err = etcdClient.Put(ctx, etcdKey, string(data))
	cancel()
	if err != nil {
		logs.Warn("Put failed, err:%v", err)
		return
	}

	logs.Debug("put etcd succ, data:%v", string(data))
	return
}

//获取数据库中所有的数据，并准备发送到Etcd中
func GetAllLogCof() {
	logConfArr = nil
	loglist, err := GetAllLogInfo()
	if err != nil {
		logs.Warn("get app list failed, err:%v", err)
		return
	}
	for _, logCof := range loglist {
		logConfArr = append(
			logConfArr,
			CollectConf{
				LogPath: logCof.LogPath,
				Topic:   logCof.Topic,
			},
		)
	}
}
