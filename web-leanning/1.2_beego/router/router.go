package router

import (
	"github.com/astaxie/beego"
	"web_application/1.2_beego/controller/IndexController"
)

func init() {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
}
