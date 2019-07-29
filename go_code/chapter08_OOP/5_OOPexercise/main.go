package main   //面向对象编程应用实例
import (
	"fmt"
)

type Student struct{        //首先定义一个结构体 字段/属性
	name string
	gender string
	age int 
	id int
	score float64

}

func (stu *Student) say()string {       //声明方法
	infostr :=fmt.Sprintf("studeng的信息 name=[%v]  gender=[%v]  age=[%v]  id=[%v]  score=[%v]",
	stu.name,stu.gender,stu.age,stu.id,stu.score)
	return infostr

}

func main(){
	var stu Student = Student{     //在main函数中创建Student的实例
		name :"teemo" ,
		gender :"男",
		age :18,
		id :20,
		score :66,
	}

	fmt.Println(stu.say())    //通过实例调用方法



}