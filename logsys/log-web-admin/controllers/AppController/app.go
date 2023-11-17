package AppController

import (
	"fmt"
	"strings"
	"time"
	"logsys/log-web-admin/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AppController struct {
	beego.Controller
}

//列表页面
func (p *AppController) AppList() { //列表页面
	logs.Debug("enter index controller")
	p.Layout = "layout/layout.html" //指导布局文件 ,布局页面整体改变  html中要在指定位置添加{{.LayoutContent}}
	appList, err := models.GetAllAppInfo()
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
		p.TplName = "app/error.html"
		logs.Warn("get app list failed, err:%v", err)
		return
	}

	logs.Debug("get app list succ, data:%v", appList)

	p.Data["applist"] = appList
	p.TplName = "app/index.html"

}

//项目申请页面
func (p *AppController) AppApply() {
	logs.Debug("enter index controller")
	p.Layout = "layout/layout.html" //指导布局文件 ,布局页面整体改变 html中要在指定位置添加{{.LayoutContent}}
	p.TplName = "app/apply.html"
}

//

func (p *AppController) DeleteList() { //列表页面
	logs.Debug("enter index controller")
	p.Layout = "layout/layout.html" //指导布局文件 ,布局页面整体改变  html中要在指定位置添加{{.LayoutContent}}
	appList, err := models.GetAllAppInfo()
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
		p.TplName = "app/error.html"
		logs.Warn("get app list failed, err:%v", err)
		return
	}

	logs.Debug("get app list succ, data:%v", appList)

	p.Data["applist"] = appList
	p.TplName = "app/delete.html"

}

// func (p *AppController) DeleteList() {

// 	logs.Debug("enter app DeleteList index controller")

// 	p.Layout = "layout/layout.html"
// 	appList, err := models.GetAllAppInfo() // 获取数据库全部日志内容，准备显示到网页上
// 	if err != nil {
// 		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
// 		p.TplName = "app/error.html"

// 		logs.Warn("get app list failed, err:%v", err)
// 		return
// 	}

// 	logs.Debug("get app list succ, data:%v", appList)
// 	p.Data["applist"] = appList

// 	p.TplName = "app/delete.html"
// }

//项目申请处理，并跳转到列表页面
func (p *AppController) AppCreate() {

	logs.Debug("enter index controller")

	appName := p.GetString("app_name") //从页面form提交的数据中获取相关内容
	appType := p.GetString("app_type")
	developPath := p.GetString("develop_path")
	ipstr := p.GetString("iplist")

	p.Layout = "layout/layout.html" //指导布局文件 ,布局页面整体改变 html中要在指定位置添加{{.LayoutContent}}
	if len(appName) == 0 || len(appType) == 0 || len(developPath) == 0 || len(ipstr) == 0 {
		p.Data["Error"] = fmt.Sprintf("非法参数")
		p.TplName = "app/error.html"

		logs.Warn("invalid parameter")
		return
	}

	appInfo := &models.AppInfo{}
	appInfo.AppName = appName
	appInfo.AppType = appType
	appInfo.DevelopPath = developPath
	appInfo.IP = strings.Split(ipstr, ",")
	appInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05") //2006-01-02 15:04:05为日期格式化固定数值

	err := models.CreateApp(appInfo) //将相关数据上传到数据库中
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("创建项目失败，数据库繁忙")
		p.TplName = "app/error.html"

		logs.Warn("invalid parameter")
		return
	}
	p.Redirect("/app/list", 302) //302重定向
}

//日志申请处理，并跳转到列表页面
func (p *AppController) AppUpdate() {
	logs.Debug("enter update controller")
	appId, err := p.GetInt("app_id")
	appName := p.GetString("app_name")
	appPath := p.GetString("app_path")
	appType := p.GetString("app_type")

	p.Layout = "layout/layout.html"
	if err != nil || len(appName) == 0 || len(appPath) == 0 || len(appType) == 0 {
		p.Data["Error"] = fmt.Sprintf("非法参数")
		p.TplName = "app/error.html"

		logs.Warn("invalid parameter")
		return
	}

	appInfo := &models.AppInfo{}
	appInfo.AppId = appId
	appInfo.AppName = appName
	appInfo.AppType = appType
	appInfo.DevelopPath = appPath
	appInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05") //2006-01-02 15:04:05为固定数值

	err = models.UpdateApp(appInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("修改项目失败，数据库繁忙")
		p.TplName = "app/error.html"

		logs.Warn("invalid parameter")
		return
	}
	p.Redirect("/app/list", 302)
}

func (p *AppController) AppDelete() {
	logs.Debug("enter delete controller")
	appId, err := p.GetInt("app_id")
	p.Layout = "layout/layout.html"
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("服务器错误，删除失败")
		p.TplName = "app/error.html"

		logs.Warn("invalid parameter")
		return
	}

	appInfo := &models.AppInfo{}
	appInfo.AppId = appId

	err = models.DeleteApp(appInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("修改项目失败，数据库繁忙")
		p.TplName = "app/error.html"

		logs.Warn("invalid parameter")
		return
	}

	p.Redirect("/app/list", 302)
}
