package main
import (
	"fmt"
)

func main(){
	var num int = 10
	var ptr *int   //定义指针变量ptr
	ptr = &num     //取num的地址赋给ptr
	fmt.Printf("num的地址是%v\n",ptr)

	*ptr = 20     //修改ptr指向地址内的值
	fmt.Printf("num的值是%v\n",num)



}