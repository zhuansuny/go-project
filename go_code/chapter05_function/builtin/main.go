package main
//内置函数
import (
	"fmt"
)

func main(){
	num1 := 100
	fmt.Printf("类型%T 值是%d 地址是%v\n",num1,num1,&num1)

	num2 := new(int)    //new 用来分配内存 ，返回的是一个指针，主要用来分配值类型（int、flaot等）
	*num2 = 100         //make 也是用来分配内存，用于分配引用类型（map等）
	fmt.Printf("num2类型%T 值是%d 地址是%v 指向的值是%v\n",num2,num2,&num2,*num2)


}