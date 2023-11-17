package process2

//处理和用户相关的请求，登陆、注册、注销、用户列表管理
import (
	"encoding/json"
	"fmt"
	"go_code/chapter16_chatRoom/common/message"
	"go_code/chapter16_chatRoom/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

//serverProcessLogin函数，专门处理登陆请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

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
	//创建utils实例
	var tran = utils.Transfer{
		Conn: this.Conn,
	}
	err = tran.WritePkg(data)
	return
}
