package main
//封装
//面向对象三大特性 ：封装、继承、多态
import (
	"fmt"
)
type person struct{       //封装就是建立一个相当于构造方法的方法，并将字段使用set get 方法封装
	Name string           //golang并没有特别强调封装，对面向对象的特性做了简化
	age int               //小写其他包不可直接访问
	salary float64

}


func NewPerson(name string) *person{  //使用一个函数创建一个person结构体实例，并将地址返回
	return &person{                        //等同于JAVA中的构造方法
		Name : name ,
	}

}

func (p *person) GetAge() int{   //age的get set方法   
	return p.age
}

func (p *person) SetAge(age int){  //可以在set 方法内对年龄加限制调整
	if age >0 && age<150 {
		p.age =age
	}else{
		fmt.Println("输入年龄范围不正确")
	}
}

func (p *person) GetSalary() int{   //salary的get set方法
	return p.age
}

func (p *person) SetSalary(salary float64) {  
	if salary >2000 && salary<1500000 {
		p.salary =salary
	}else{
		fmt.Println("输入工资范围不正确")
	}
}

func main(){
	p :=person{
		Name :"jack",
	}

	p.SetAge(1520)  //年龄不合理，报错
	p.SetAge(18)	 //正常
	fmt.Println(p.GetAge()) //输出为18

}