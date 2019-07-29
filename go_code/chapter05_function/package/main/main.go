package main   //main包只能有一个
import (
	"go_code/chapter05_function/package/utils"//导入utils包(默认从$GOPATH的路径的src开始)
)
func main(){
	utils.Cal()//可以正常调用utils包中的Cal函数（函数名首字母必须是大写） 
	utils.Cal2()
}