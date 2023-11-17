package models //mysql数据库相关操作

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/astaxie/beego/logs"
)

type AppInfo struct {
	AppId       int      `db:"app_id"`
	AppName     string   `db:"app_name"`
	AppType     string   `db:"app_type"`
	CreateTime  string   `db:"create_time"`
	DevelopPath string   `db:"develop_path"`
	IP          []string `db:"ip"`
}

var (
	Db *sqlx.DB
)

func InitDb(db *sqlx.DB) { //初始数据库
	Db = db
}

//----------获取数据库表格tbl_app_info中的数据----------------

func GetAllAppInfo() (applist []AppInfo, err error) {
	err = Db.Select(&applist, "select app_id, app_name, app_type, create_time, develop_path from tbl_app_info ")

	if err != nil {
		logs.Warn("Get All App Info failed, err:%v", err)
		fmt.Println("Get All App Info failed, err:", err)
		return
	}
	return
}

func GetIPInfoById(appId int) (iplist []string, err error) {
	err = Db.Select(&iplist, "select ip from tbl_app_ip where app_id=?", appId)
	if err != nil {
		logs.Warn("Get app_id _IP Info failed, err:%v", err)
		return
	}
	return

}

func GetIPInfoByName(appName string) (iplist []string, err error) {

	var appId []int
	err = Db.Select(&appId, "select app_id from tbl_app_info where app_name=?", appName)
	if err != nil || len(appId) == 0 {
		logs.Warn("select app_id failed, Db.Exec error:%v", err)
		return
	}

	err = Db.Select(&iplist, "select ip from tbl_app_ip where app_id=?", appId[0])
	if err != nil {
		logs.Warn("Get app_name _IP Info failed, err:%v", err)
		return
	}
	return
}

// ---- 更新项目表格信息--------

func UpdateApp(info *AppInfo) (err error) {

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

	r, err := conn.Exec("update tbl_app_info set develop_path=?, app_name=?,create_time=?, app_type=? where app_id =? ",
		info.DevelopPath, info.AppName, info.CreateTime, info.AppType, info.AppId)

	if err != nil {
		logs.Warn("UpdateApp failed, Db.Exec error:%v", err)
		return
	}

	_, err = r.LastInsertId()
	if err != nil {
		logs.Warn("UpdateLog failed, Db.LastInsertId error:%v", err)
		return
	}

	return
}

//----添加表格信息---------------

func CreateApp(info *AppInfo) (err error) {
	conn, err := Db.Begin() //开启事务
	if err != nil {
		logs.Warn("createApp falied, Db.Begin error:%v", err)
		return
	}
	defer func() {
		if err != nil {
			conn.Rollback()
			return
		}
		conn.Commit()
	}()

	r, err := conn.Exec("insert into tbl_app_info(app_name, app_type, develop_path, create_time)values(?, ?, ?, ?)",
		info.AppName, info.AppType, info.DevelopPath, info.CreateTime)

	if err != nil {
		logs.Warn("CreateApp failed, Db.Exec error:%v", err)
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		logs.Warn("CreateApp failed, Db.LastInsertId error:%v", err)
		return
	}

	for _, ip := range info.IP {
		_, err = conn.Exec("insert into tbl_app_ip(app_id, IP)values(?,?)", id, ip)
		if err != nil {
			logs.Warn("CreateAppIp failed, conn.Exec ip error:%v", err)
			return
		}
	}
	return

}

//删除数据数据库数据
func DeleteApp(info *AppInfo) (err error) {

	r, err := Db.Exec("delete from tbl_app_info where app_id=? ", info.AppId)

	if err != nil {
		logs.Warn("delete app failed, Db.Exec error:%v", err)
		return
	}

	_, err = r.LastInsertId()
	if err != nil {
		logs.Warn("delete app failed, Db.LastInsertId error:%v", err)
		return
	}

	return
}
