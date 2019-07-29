package main  //方法，通常用于结构体
import (
	"fmt"
)

type Person struct {
	Name string       //字段使用大写，在别的包可以调用
}

func (p Person) test(){          //(p Person)的意思是只能由Person结构体的实例来调用这个方法  
	p.Name = "小红"               
	fmt.Println("test()",p.Name)  //输出"小红""
}

func (p Person) speak(){
	fmt.Println(p.Name,"是一个好人")
}

func (p Person) sum(){
	res :=0
	for i := 0; i < 1000; i++ {
		res +=i
	}
	fmt.Println(p.Name,"计算结果是",res)
}

func (p Person) sum2(n int){
	res :=0
	for i := 0; i < n; i++ {
		res +=i
	}
	fmt.Println(p.Name,"计算结果是",res)
}

func (p Person) sum3(n int,m int)int{
	return n+m
}

func main(){
	var p Person
	p.Name  = "小明"
	p.test()      //调用的同时将p也当成实参传给test
	p.speak()
	p.sum()
	p.sum2(100)  ////调用的同时将p和100当成实参传给sum2
	fmt.Println(p.sum3(100,200))
	fmt.Println(p.Name)   //由于结构体是值传递，输出还是“小明”
}