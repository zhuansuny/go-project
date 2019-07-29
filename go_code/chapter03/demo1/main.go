package main
import "fmt"

func main(){
	var b int //int 初始值为0
	fmt.Println("b=",b)
	var a = 11//可以根据值进行类型推导
	fmt.Println("a=",a)
	name := "tom"//等价于 var name string ="tom"
	fmt.Println("name=",name)
	//go支持多变量声明
	var v1,v2 int = 1,2//同时声明同类型多个变量
	fmt.Println("V1=",v1,"V2=",v2)
	var n1, name1, n3 =100, "tom", 888//同时声明不同同类型多个变量
	fmt.Println("n1=",n1,"name1=",name1,"n3=",n3)
}