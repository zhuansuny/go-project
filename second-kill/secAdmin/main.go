package main

import (
	_ "secKill/secAdmin/routers"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

func main() {
	err := initAll()
	if err != nil {
		logs.Error("init falied, err:", err)
	}
	beego.Run()
}
