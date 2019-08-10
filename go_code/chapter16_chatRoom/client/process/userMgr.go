package process //维护用户在线列表

import (
	"fmt"
	"go_code/chapter16_chatRoom/client/model"
	"go_code/chapter16_chatRoom/common/message"
)

//客户端要维护的map  全局变量
var onlineUsers map[int]*message.User = make(map[int]*message.User, 100)
var CurUser model.CurUser //我们在用户登录成功后，完成对CurUser初始化

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserOnlineStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	fmt.Printf("用户:%d上线\n", notifyUserStatusMes.UserId)
}

//用户离线
func updateUserOfflineStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	delete(onlineUsers, notifyUserStatusMes.UserId)

	//在客户端显示离线的用户
	fmt.Printf("用户:%d离线\n", notifyUserStatusMes.UserId)
}

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历一把 onlineUsers
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		//如果不显示自己.
		fmt.Println("用户id:\t", id)
	}
}
