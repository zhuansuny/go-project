package main
import (
	"fmt"
)
//init函数，每一个源文件都可以包含一个init函数，可以在main函数之前执行（全局变量最先执行）
//通常在init函数中进行初始化

func init(){                    //init通常起初始化作用
	fmt.Println("init()...",age)
}

var age =90  //全局变量最先执行

func main(){
	fmt.Println("main()...")
}