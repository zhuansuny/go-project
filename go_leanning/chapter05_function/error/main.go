package main
//golang错误处理机制
import (
	"fmt"
	"errors"
)

func test(){
	defer func(){
		err := recover()              //捕获错误，程序可以继续运行
		if err != nil{
			fmt.Println("err=",err)  
		}
	}()
	num := 10
	num1 := 0
	res := num/num1          //不能除于0，程序会报错并退出（崩溃）
	fmt.Println("res=",res)
}

//函数去读取配置文件的信息
//如果名称不正确，我们返回一个自定义的错误
func readCnfig(name string)(err error){
	if name == "config.ini" {
		return nil
	}else{
		return errors.New("读取配置文件失败")
	}
}

func test2(){
	err := readCnfig("config2.ini")
	if err != nil{
		panic(err)  //如果读取文件错误，就输出这个错误，并终止程序
	}
	fmt.Println("test2继续执行")
}

func main(){
	test()
	test2()

}