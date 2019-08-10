package message

//定义一个User结构体
type User struct {
	//确定字段信息
	//为了序列化和反序列化成功，
	//用户信息的json字符串的key 和 结构体的字段对应的 tag 名字一致!!!
	UserId     int    `json:userId`
	UserName   string `json:userName`
	UserPwd    string `json:userPwd`
	UserStatus int    `json:userStatus`
}
