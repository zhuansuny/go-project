package process2 //维护用户在线列表

import (
	"fmt"
)

//定义一个UserMgr结构体的实例
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr的初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//添加onlineUser
func (this *UserMgr) AddonlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除onlineUser
func (this *UserMgr) DelonlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线的用户
func (this *UserMgr) GetAllonlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetonlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok { //查找的这个用户，当前不在线。
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return

}
