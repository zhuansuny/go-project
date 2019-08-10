package process

//显示登陆界面
//保持和服务器的通讯（即启动一个协程）
//当读取服务器发送的消息后，就会显示在界面
import (
	"encoding/json"
	"fmt"
	"go_code/chapter16_chatRoom/client/utils"
	"go_code/chapter16_chatRoom/common/message"
	"net"
	"os"
)

func ShowMenu(userId int) {
	fmt.Printf("---------恭喜用户%d登录成功-----------\n", userId)
	fmt.Println("---------1.显示在线用户列表-----------")
	fmt.Println("---------2.发送群聊消息------------------")
	fmt.Println("---------3.发送私聊消息------------------")
	fmt.Println("---------4.信息列表-----------")
	fmt.Println("---------5.退出系统-----------")
	fmt.Println("请选择（1-4）：")
	var key int
	var content string
	var Id int
	fmt.Scan(&key)
	//因为经常使用SmsProcess，因此我们将其定义在switch外面
	smsProcess := &SmsProcess{}
	switch key {
	case 1:
		outputOnlineUser()
		//fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("请输入想要输入的消息：")
		for {
			fmt.Scan(&content)
			if content == "exit" {
				break
			}
			smsProcess.SentGroupMes(content)
		}
	case 3:
		fmt.Println("请输入想要聊天用户的Id：")
		fmt.Scan(&Id)
		fmt.Println("请输入想要输入的消息：")
		for {
			fmt.Scan(&content)
			if content == "exit" {
				break
			}
			smsProcess.SentPrivateMes(Id, content)
		}
	case 4:
		fmt.Println("信息列表")
	case 5:
		fmt.Println("退出系统")
		smsProcess.OfflineMes(userId)
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
		//fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg() //客户端不断的读取，会堵塞在这里
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//如果读取到消息，进行下一步处理逻辑
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//1. 取出.NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("json 反序列化失败2 err=", err)
				return
			}
			if notifyUserStatusMes.Status == 0 {

				//var notifyUserStatusMes message.NotifyUserStatusMes
				//json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
				//2. 把这个用户的信息，状态保存到客户map[int]User中
				updateUserOnlineStatus(&notifyUserStatusMes)
			} else {
				updateUserOfflineStatus(&notifyUserStatusMes)
			}

		case message.SmsMesType:
			outputGrounp(&mes)
		case message.SmMesType:
			outputGrounp(&mes)
		default:
			fmt.Println("读取了一个无法识别的指令")
		}
	}
}
