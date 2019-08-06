package main

import (
	"encoding/json"
	"fmt"
	"go_code/chapter16_chatRoom/commend/message"
	"net"
	_ "time"
)

func login(userId int, Password string) (err error) {
	//fmt.Printf("userID = %d Password = %s\n", userID, Password)
	//return nil

	//1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != err {
		fmt.Println("链接失败，err=", err)
		return
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建一个LoginMes结构体

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = Password

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != err {
		fmt.Println("json序列化失败，err=", err)
		return
	}
	//5.data赋值给mes.Data字段
	mes.Data = string(data)
	fmt.Println("type", mes.Type)
	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != err {
		fmt.Println("json序列化失败，err=", err)
		return
	}

	//7.data就是要发送的数据
	//7.1先将data的长度发送给服务器
	//先获取到data的长度，转换成一个表示长度的byte切片
	writePkg(conn, data)

	// time.Sleep(10 * time.Second)
	// fmt.Println("休眠了20秒")
	mes, err = readPkg(conn)

	if err != nil {
		fmt.Println("conn.ReadPkg(conn) fail", err)
		return
	}
	//将mes的Data的反序列化成 loginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("反序列化3 fail", err)
		return
	}
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return
}
