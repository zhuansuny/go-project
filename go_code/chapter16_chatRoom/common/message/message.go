package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"` //消息的类型
	Data string `json:"data"` //消息的内容
}

type LoginMes struct {
	UserId   int    `json:"userId"`   //用户ID
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

type LoginResMes struct {
	Code  int    `json:"code"` //返回状态码500表示用户未注册，200表示登陆成功
	Error string `json:"error"`
}

type RegisterMes struct {
}
