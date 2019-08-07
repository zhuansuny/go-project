package process

import (
	"fmt"
	"go_code/chapter16_chatRoom/client/utils"
	"net"
	"os"
)

//显示登陆界面
//保持和服务器的通讯（即启动一个协程）
//当读取服务器发送的消息后，就会显示在界面
func ShowMenu() {
	fmt.Println("---------恭喜xxx登录成功-----------")
	fmt.Println("---------1.显示在线用户列表-----------")
	fmt.Println("---------2.发送消息------------------")
	fmt.Println("---------3.信息列表-----------")
	fmt.Println("---------4.退出系统-----------")
	fmt.Println("请选择（1-4）：")
	var key int
	fmt.Scan(&key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入数字有误，请重新输入")
	}
}

//和服务器端保持连接
func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例，不同的读取服务器发送的消息
	tf := utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg() //客户端不断的读取，会堵塞在这里
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//如果读取到消息，进行下一步处理逻辑
		fmt.Println("mes=", mes)
	}
}
