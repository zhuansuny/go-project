package process2 //维护用户离线列表

import (
	"go_code/chapter16_chatRoom/common/message"
)

//定义一个UserOfflineMes结构体的实例
var (
	userOfflineMes *UserOfflineMes
)

type UserOfflineMes struct {
	offlineUsers map[int]*message.SmMes
}

//完成对UserOfflineMes的初始化工作
func init() {
	userOfflineMes = &UserOfflineMes{
		offlineUsers: make(map[int]*message.SmMes, 1024),
	}
}

//添加offlineUsers
func (this *UserOfflineMes) AddofflineUser(userId int, smMes message.SmMes) {
	this.offlineUsers[userId] = &smMes
}

//删除offlineUsers
func (this *UserOfflineMes) DelonlineUser(userId int) {
	delete(this.offlineUsers, userId)
}

// //返回当前所有在线的用户
// func (this *UserOfflineMes) GetAllonlineUser() map[int]*UserProcess {
// 	return this.onlineUsers
// }

// //根据id返回对应的值
// func (this *UserOfflineMes) GetonlineUserById(userId int) (up *UserProcess, err error) {
// 	up, ok := this.onlineUsers[userId]
// 	if !ok { //查找的这个用户，当前不在线。
// 		err = fmt.Errorf("用户%d不存在", userId)
// 		return
// 	}
// 	return

// }
