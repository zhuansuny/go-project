package main //总的处理器，根据客户端的请求，调用相应的处理器，完成相应的任务

import (
	"fmt"
	"go_code/chapter16_chatRoom/common/message"
	"go_code/chapter16_chatRoom/server/process"
	"go_code/chapter16_chatRoom/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//根据客户端发送的消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	//fmt.Println("mes=", mes)
	switch mes.Type {
	case message.LoginMesType:
		//处理登陆

		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)

	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		sp := &process2.SysProcess{}

		err = sp.SendGroundMes(mes)
	case message.SmMesType:
		sp := &process2.SysProcess{}

		err = sp.SentPrivateMes(mes)
	case message.NotifyUserStatusMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.NotifyOthersOfflineUser(mes)

	default:
		fmt.Println("消息类型错误，无法处理")
	}
	return
}

func (this *Processor) process2() (err error) {

	//循环的客户端发送的信息
	for {
		//读取数据包，直接封装成一个函数readPkg(), 返回Message, Err
		//创建一个Transfer 实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出..")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}

		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}

}
