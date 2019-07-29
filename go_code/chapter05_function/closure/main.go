package main  //闭包
import (
	"fmt"
	"strings"
)

func AddUpper() func (int) int {  //返回的是匿名函数
	var n int = 10
	return func(x int) int {     
		n = n + x                  //匿名函数引用函数外的n,
		return n
	}
}

func makeSuffix(suffix string) func(string) string{   //相当于JAVA构造方法，在赋值时需要确定suffix的值
	return func(name string) string{
		if !strings.HasSuffix(name ,suffix){
			return name +suffix
		}
		return name
	}
}

func main(){
	f := AddUpper()        //函数赋值给f时会初始化一次（相当于JAVA建立对象），之后的n值会变化
	fmt.Println(f(1))      //n等于初始的10，结果为11
	fmt.Println(f(2))      //n等于11，结果为13
	fmt.Println(f(3))      //n等于13，结果为16

	g := AddUpper()        //函数重新赋值给g时,会初始化n=10
	fmt.Println(g(1))

	h := makeSuffix(".jpg")
	fmt.Println(h("winter"))
}