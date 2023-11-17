package main
//工厂模式
import (
	"fmt"
	"go_code/chapter08_OOP/6_factory/model"
)


func main(){
	var stu = model.Student{   //大写的结构体在别的包可以调用
		Name : "tom",        //大写的字段在别的包可以调用
		Score : 78.9,
	}
//如果是小写的可以使用工厂模式进行调用
	                         
	fmt.Println(stu)

	stu1 := model. NewStudent("jack",88.1) //x=结构体小写可以使用工厂模式间接调用
	//fmt.Println(stu1.name) //但小写的字段无法访问
	stu1.SetName("teemo")    //可以通过get set方法进行修改调用
	fmt.Println(stu1.GetName())  
}