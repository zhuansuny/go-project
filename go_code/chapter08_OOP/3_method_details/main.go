package main  //方法，通常用于结构体
import (
	"fmt"
)

type integer int

func (i integer) print(){  //方法可以用于其他自定义的类型
	fmt.Println("i=",i)
}

func (i *integer) change(){ //传入指针，修改数据
	*i = *i + 1
}


type student struct{
	name string
	age int
}

func (stu *student) String() string{     //toString方法  引用传递相对于值传递速度快
	str := fmt.Sprintf("name=[%v] age=[%v] ",stu.name,stu.age)
	return str
}

func main(){
	var i integer = 30
	i.print()
	i.change()
	fmt.Println("i=",i)
	var stu student
	stu.name = "mary"
	stu.age = 18
	fmt.Println(&stu)  //




}