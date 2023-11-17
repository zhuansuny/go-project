package main

import (
	"github.com/astaxie/beego"
	_ "web_application/1.2_beego/router"
	//需要在gopath目录下cmd 运行 go get github.com/astaxie/beego 安装go的beego包
)

// type HomeController struct {
// 	beego.Controller
// }

// func (this *HomeController) GET() {
// 	this.Ctx.WriteString("hello world")
// }

func main() {
	//beego.Router("/", &HomeController{}, "*:GET")
	beego.Run()
}
