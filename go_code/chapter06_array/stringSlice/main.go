package main  //字符串切片
import (
	"fmt"
)
func main(){
	str := "中helloWorld" //字符串底层是byte，可以使用切片
	slice := str[5:]
	fmt.Println(slice)
	//str[0]='z'   //string不可变，也就是不能通过str[0]直接赋值来进行修改
	arr :=[]byte(str)  //可以先转成byte数组修改，再转成字符串
	arr1 :=[]rune(str)  //转成runes切片，可以解决中文问题
	arr[0] ='z'
	arr1[0] ='z'
	str = string(arr1)
	fmt.Println(str)
	str = string(arr)
	fmt.Println(str)
}