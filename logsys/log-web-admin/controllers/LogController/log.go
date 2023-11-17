package LogController

import (
	"fmt"
	model "logsys/log-web-admin/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LogController struct {
	beego.Controller
}

//列表页面
func (p *LogController) LogList() {

	logs.Debug("enter index controller")

	p.Layout = "layout/layout.html"
	logList, err := model.GetAllLogInfo() // 获取数据库全部日志内容，准备显示到网页上
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
		p.TplName = "app/error.html"

		logs.Warn("get app list failed, err:%v", err)
		return
	}

	logs.Debug("get app list succ, data:%v", logList)
	p.Data["loglist"] = logList

	p.TplName = "log/index.html"
}

//日志创建页面
func (p *LogController) LogApply() {

	logs.Debug("enter index controller")
	p.Layout = "layout/layout.html"
	p.TplName = "log/apply.html"
}

//
func (p *LogController) DeleteList() {

	logs.Debug("enter index controller")

	p.Layout = "layout/layout.html"
	logList, err := model.GetAllLogInfo() // 获取数据库全部日志内容，准备显示到网页上
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
		p.TplName = "app/error.html"

		logs.Warn("get app list failed, err:%v", err)
		return
	}

	logs.Debug("get app list succ, data:%v", logList)
	p.Data["loglist"] = logList

	p.TplName = "log/delete.html"
}

//日志申请处理，并跳转到列表页面
func (p *LogController) LogCreate() {

	logs.Debug("enter index controller")
	appName := p.GetString("app_name")
	logPath := p.GetString("log_path")
	topic := p.GetString("topic")

	p.Layout = "layout/layout.html"
	if len(appName) == 0 || len(logPath) == 0 || len(topic) == 0 {
		p.Data["Error"] = fmt.Sprintf("非法参数")
		p.TplName = "log/error.html"

		logs.Warn("invalid parameter")
		return
	}

	logInfo := &model.LogInfo{}
	logInfo.AppName = appName
	logInfo.LogPath = logPath
	logInfo.Topic = topic
	logInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05") //2006-01-02 15:04:05为固定数值

	err := model.CreateLog(logInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("创建项目失败，数据库繁忙")
		p.TplName = "log/error.html"

		logs.Warn("invalid parameter")
		return
	}

	//iplist, err := model.GetIPInfoByName(appName)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("获取项目ip失败，数据库繁忙")
		p.TplName = "log/error.html"

		logs.Warn("invalid parameter")
		return
	}
	//keyFormat := "logagent/config/%s"

	//for _, ip := range iplist {
	//key := fmt.Sprintf(keyFormat, ip)
	key := "/logAgent/config/127.0.0.1"
	err = model.SetLogConfToEtcd(key)
	if err != nil {
		logs.Warn("Set log conf to etcd failed, err:%v", err)
		return
	}
	//}
	p.Redirect("/log/list", 302)
}

//日志更新处理
func (p *LogController) LogUpdate() {
	logs.Debug("enter update controller")
	logId, err := p.GetInt("log_id")
	appName := p.GetString("app_name")
	logPath := p.GetString("log_path")
	topic := p.GetString("topic")

	p.Layout = "layout/layout.html"
	if err != nil || len(appName) == 0 || len(logPath) == 0 || len(topic) == 0 {
		p.Data["Error"] = fmt.Sprintf("非法参数")
		p.TplName = "log/error.html"

		logs.Warn("invalid parameter")
		return
	}

	logInfo := &model.LogInfo{}
	logInfo.LogId = logId
	logInfo.AppName = appName
	logInfo.LogPath = logPath
	logInfo.Topic = topic
	logInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05") //2006-01-02 15:04:05为固定数值

	err = model.UpdateLog(logInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("修改项目失败，数据库繁忙")
		p.TplName = "log/error.html"

		logs.Warn("invalid parameter")
		return
	}

	key := "/logAgent/config/127.0.0.1"
	err = model.SetLogConfToEtcd(key)
	if err != nil {
		logs.Warn("Set log conf to etcd failed, err:%v", err)
		return
	}
	//}
	p.Redirect("/log/list", 302)
}

//日志删除处理，并跳转到列表界面
func (p *LogController) LogDelete() {
	logs.Debug("enter delete controller")
	logId, err := p.GetInt("log_id")
	p.Layout = "layout/layout.html"
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("服务器错误，删除失败")
		p.TplName = "log/error.html"

		logs.Warn("invalid parameter")
		return
	}

	logInfo := &model.LogInfo{}
	logInfo.LogId = logId

	err = model.DeleteLog(logInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("修改项目失败，数据库繁忙")
		p.TplName = "log/error.html"

		logs.Warn("invalid parameter")
		return
	}

	key := "/logAgent/config/127.0.0.1"
	err = model.SetLogConfToEtcd(key)
	if err != nil {
		logs.Warn("Set log conf to etcd failed, err:%v", err)
		return
	}
	//}
	p.Redirect("/log/list", 302)
}
