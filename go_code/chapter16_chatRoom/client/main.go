package main

import (
	"fmt"
)

func loginSurface() {
	var userId int
	var Password string
	fmt.Println("请输入用户的id")
	fmt.Scanf("%d\n", &userId)

	fmt.Println("请输入用户的密码")
	fmt.Scanf("%v\n", &Password)
	err := login(userId, Password)
	if err != nil {
		fmt.Println("登陆失败，err=", err)
	}
}

func registSurface() {
	var userId int
	var Password string
	fmt.Println("请输入用户的id")
	fmt.Scanf("%d\n", &userId)

	fmt.Println("请输入用户的密码")
	fmt.Scanf("%v\n", &Password)
	err := login(userId, Password)
	if err != nil {
		fmt.Println("登陆失败，err=", err)
	}
}

func main() {
	var key int
	var loop = true
	for loop {
		fmt.Println("---------------------------欢迎登陆多人聊天系统---------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loginSurface()
			loop = false
		case 2:
			fmt.Println("注册用户")
			registSurface()
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入的数字有误，请重新输入")

		}
	}

}
