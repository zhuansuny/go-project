package main

import (
	"encoding/json"
	"fmt"
	"go_code/chapter16_chatRoom/commend/message"
	"io"
	"net"
)

//根据客户端发送的消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登陆
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:

	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

//serverProcessLogin函数，专门处理登陆请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {

	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json 反序列化失败2 err=", err)
		return
	}
	//1先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//2再声明一个 LoginResMes，并完成赋值
	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用"
	}
	//3将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//4. 将data 赋值给 resMes
	resMes.Data = string(data)

	//5. 对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	err = writePkg(conn, data)
	return
}

func process(conn net.Conn) {
	//需要延时关闭conn
	defer conn.Close()

	//循环读取客户端发送信息

	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("readPkg fail err=", err)
			return
		}

		fmt.Println("mes=", mes.Data)
		err = serverProcessMes(conn, &mes)
		if err == nil {
			return
		}
	}
}

func main() {
	fmt.Println("服务器在8889端口监听。。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")

	defer listen.Close()
	if err != err {
		fmt.Println("监听失败，err=", err)
		return
	}

	for {
		fmt.Println("等待客户端来链接服务器。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接受失败，err=", err)
		}

		//一旦链接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}

}
