package routers

import (
	"secKill/secProxy/controllers"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

func init() {
	logs.Debug("enter router init")
	beego.Router("/seckill", &controllers.SkillController{}, "*:SecKill")
	beego.Router("/secinfo", &controllers.SkillController{}, "*:SecInfo")
}
