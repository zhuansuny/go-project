package main

// 1：客户端长短连接的管理
// 2：请求数据的合法性，正确性校验。
// 3：cookie管理
// 4：将请求转发到逻辑层
// 5：建立与客户端通信的加密通道。
// 6：安全防御
// 7：整合内部少量的长连接

import (
	_ "secKill/secProxy/routers"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

func main() {
	err := initConfig()
	if err != nil {
		panic(err)
		return
	}

	err = initSec()
	if err != nil {
		logs.Error("initSec failed, err:", err)
		return
	}
	beego.Run()
}
