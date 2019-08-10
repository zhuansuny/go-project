package process2

//处理和用户相关的请求，登陆、注册、注销、用户列表管理
import (
	"encoding/json"
	"fmt"
	"go_code/chapter16_chatRoom/common/message"
	"go_code/chapter16_chatRoom/server/model"
	"go_code/chapter16_chatRoom/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

//用户状态信息
func (this *UserProcess) NotifyOthersOfflineUser(mes *message.Message) (err error) {
	//遍历 onlineUsers,然后一个一个的发送NotifyUserStatus
	var notifyUserStatusMes message.NotifyUserStatusMes
	err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
	if err != nil {
		fmt.Println("json 反序列化失败2 err=", err)
		return
	}
	this.UserId = notifyUserStatusMes.UserId
	//将用户从在线列表中删除
	userMgr.DelonlineUser(this.UserId)
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == this.UserId {
			continue
		}
		up.NotifyMeOffline(this.UserId)
	}
	return
}

func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历 onlineUsers,然后一个一个的发送NotifyUserStatus
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装NotifyUserStatus
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给 mes.Data
	mes.Data = string(data)
	//对mes再次序列化，准备发送.
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送，创建Transfer实例
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}

}

func (this *UserProcess) NotifyMeOffline(userId int) {
	//组装NotifyUserStatus
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOffline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给 mes.Data
	mes.Data = string(data)
	//对mes再次序列化，准备发送.
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送，创建Transfer实例
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}

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

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}

	} else {
		loginResMes.Code = 200
		//登陆成功，将该用户添加到userMgr中
		//将登陆成功的userId赋给this
		this.UserId = loginMes.UserId
		userMgr.AddonlineUser(this)
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前在线用户的id放入到loginResMes.UsersId
		//遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "登陆成功")
	}

	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	loginResMes.Code = 200
	// } else {
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "该用户不存在，请注册再使用"
	// }
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

//serverProcessRegister函数，专门处理注册请求
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json 反序列化失败7 err=", err)
		return
	}
	//1先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//2再声明一个 RegisterResMes，并完成赋值
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			registerResMes.Code = 500
			registerResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			registerResMes.Code = 403
			registerResMes.Error = err.Error()
		} else if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 300
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "服务器内部错误"
		}

	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功")
	}

	//将registerResMes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//4. 将data 赋值给 resMes.Data
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
