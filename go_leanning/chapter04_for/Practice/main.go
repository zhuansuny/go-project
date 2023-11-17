package main
import (
	"fmt"
)
//键盘输入用户名以及密码，判断是否正确
func main(){
	var password int = 1234
	var password1 int
	var userName string = "张三"
	var userName1 string 
	var count int 
	for{
		fmt.Println("请输入用户名以及密码，使用空格隔开")
		fmt.Scanf("%s %d",&userName1,&password1)
		if(count<=4){
			fmt.Println("还有",5-count,"次机会")
		}else{
			fmt.Println("请明日再来")
			break
		}
		if(password==password1&&userName==userName1){
			fmt.Println("登陆成功")
			break
		}else{
			fmt.Println("密码或用户名错误，请重新登陆")
			count++
		}
	}

}