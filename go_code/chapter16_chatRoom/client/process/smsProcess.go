package process

import (
	"encoding/json"
	"fmt"
	"go_code/chapter16_chatRoom/client/utils"
	"go_code/chapter16_chatRoom/common/message"
)

//处理和短消息相关的逻辑 私聊、群发
type SmsProcess struct {
}

//发送群聊的消息
func (this *SmsProcess) SentGroupMes(content string) (err error) {
	//创建一个mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
		return
	}
	var tf = utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}

	return

}

//发送私聊的消息
func (this *SmsProcess) SentPrivateMes(Id int, content string) (err error) {
	//创建一个mes
	var mes message.Message
	mes.Type = message.SmMesType

	//创建一个SmsMes实例
	var smMes message.SmMes
	smMes.Content = content
	smMes.UserId = CurUser.UserId
	smMes.AcceptUserId = Id

	data, err := json.Marshal(smMes)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
		return
	}
	var tf = utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}

	return

}

//发送用户离线消息
func (this *SmsProcess) OfflineMes(userId int) (err error) {
	//创建一个mes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	//创建一个SmsMes实例
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.Status = 1
	notifyUserStatusMes.UserId = userId
	//notifyUserStatusMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
		return
	}
	var tf = utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}

	return

}
