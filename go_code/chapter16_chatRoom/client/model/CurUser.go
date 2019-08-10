package model

import (
	"go_code/chapter16_chatRoom/common/message"
	"net"
)

//因为在客户端会使用到该结构体，需要设置一个全局变量
type CurUser struct {
	message.User //继承User
	Conn         net.Conn
}
