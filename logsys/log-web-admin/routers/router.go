package routers

import (
	"logsys/log-web-admin/controllers/AppController"
	"logsys/log-web-admin/controllers/LogController"

	"github.com/astaxie/beego"
)

//路由设置
func init() {
	beego.Router("/", &AppController.AppController{}, "*:AppList")
	beego.Router("/index", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/list", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/apply", &AppController.AppController{}, "*:AppApply")
	beego.Router("/app/create", &AppController.AppController{}, "*:AppCreate")
	beego.Router("/app/deletelist", &AppController.AppController{}, "*:DeleteList")
	beego.Router("/app/update", &AppController.AppController{}, "*:AppUpdate")
	beego.Router("/app/delete", &AppController.AppController{}, "*:AppDelete")

	beego.Router("/log/apply", &LogController.LogController{}, "*:LogApply")
	beego.Router("/log/list", &LogController.LogController{}, "*:LogList")
	beego.Router("/log/create", &LogController.LogController{}, "*:LogCreate")
	beego.Router("/log/deletelist", &LogController.LogController{}, "*:DeleteList")
	beego.Router("/log/update", &LogController.LogController{}, "*:LogUpdate")
	beego.Router("/log/delete", &LogController.LogController{}, "*:LogDelete")
}
