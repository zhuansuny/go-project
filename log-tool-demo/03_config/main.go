package main //读取配置文件的内容

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "./logagent.conf") //配置文件的路径
	if err != nil {
		fmt.Println("new config failed err:", err)
		return
	}
	port, err := conf.Int("server::port") //读取[server]下的port的内容
	if err != nil {
		fmt.Println("read server:port failed,err :", err)
		return
	}
	fmt.Println("port =", port)

	log_level := conf.String("logs::log_level")
	fmt.Println("log_level:", log_level)

	log_path := conf.String("logs::log_path")
	fmt.Println("log_path:", log_path)

}
