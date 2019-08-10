package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chapter16_chatRoom/common/message"
	"go_code/chapter16_chatRoom/server/utils"
	"net"
)

type SysProcess struct {
}

func (this *SysProcess) SendGroundMes(mes *message.Message) (err error) {
	//遍历服务器端的onlineUsers map[int]*UserProcess
	//将消息转发取出
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	data, err := json.Marshal(mes) //将mes序列化
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId { //将自己过滤掉
			continue
		}
		this.SendGroundOnlineUser(data, up.Conn)
	}
	return

}

func (this *SysProcess) SentPrivateMes(mes *message.Message) (err error) {
	//遍历服务器端的onlineUsers map[int]*UserProcess
	//将消息转发取出
	var smMes message.SmMes
	err = json.Unmarshal([]byte(mes.Data), &smMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	data, err := json.Marshal(mes) //将mes序列化
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id != smMes.AcceptUserId {
			continue
		}
		this.SendGroundOnlineUser(data, up.Conn)
	}
	return

}

//发送消息
func (this *SysProcess) SendGroundOnlineUser(data []byte, conn net.Conn) (err error) {
	//创建一个Tranfer
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
	}
	return
}
