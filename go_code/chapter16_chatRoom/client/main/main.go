package main

//显示一级菜单
//根据用户的输入去调用对应的处理器
import (
	"fmt"
	"go_code/chapter16_chatRoom/client/process"
)

func loginSurface() {

}

func main() {
	var key int
	var userId int
	var Password string
	for {
		fmt.Println("---------------------------欢迎登陆多人聊天系统---------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)

			fmt.Println("请输入用户的密码")
			fmt.Scanf("%v\n", &Password)
			up := &process.UserProcess{}
			err := up.Login(userId, Password)
			if err != nil {
				fmt.Println("登陆失败，err=", err)
			}

		case 2:
			fmt.Println("注册用户")
			//registSurface()
		case 3:
			fmt.Println("退出系统")
			break
		default:
			fmt.Println("输入的数字有误，请重新输入")

		}
	}

}
