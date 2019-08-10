package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	SmsMesType              = "SmsMes"
	SmMesType               = "SmMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

//定义几个用户状态常量
const (
	UserOnline     = 0 //用户在线
	UserOffline    = 1 //用户离线
	UserBusyStatus = 2 //用户忙
)

//消息结构体
type Message struct {
	Type string `json:"type"` //消息的类型
	Data string `json:"data"` //消息的内容
}

//登陆消息结构体
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户ID
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

//登陆返回消息结构体
type LoginResMes struct {
	Code    int    `json:"code"` //返回状态码500表示用户未注册，200表示登陆成功
	UsersId []int  `json:"userId"`
	Error   string `json:"error"`
}

//注册信息结构体
type RegisterMes struct {
	User User `json:"user"` //类型是User结构体
}

//注册返回消息结构体
type RegisterResMes struct {
	Code  int
	Error string
}

//为例配合服务器推送用户状态变化的消息

type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户状态
}

//群聊消息结构体
type SmsMes struct {
	Content string `json:"content"`
	User           //继承User结构体
}

//私聊消息结构体
type SmMes struct {
	Content      string `json:"content"`
	User                //继承User结构体
	AcceptUserId int
	Flag         int //判断是否发送成功 0表示发送成功 ，1表示用户离线，2表示用户不存在
}
