package main   //面向对象三大特性之继承
import (
	"fmt"
	"go_code/chapter09_OOPpart2/3_extends/student"
)

// type Student struct {   
// 	Name string
// 	Age int
// 	Score int
// }

// func (p *Student) ShowInfo(){ 
// 	fmt.Printf("学生%v 年龄=%v 成绩=%v\n",p.Name,p.Age,p.Score)
// }

// func (p *Student) SetScore(score int ){
// 	p.Score = score
	
// }
// func (p *Student) testing(){
// 	fmt.Printf("学生%v 正在考试\n",p.Name)
	
// }
type Stu student.Student 

type Pupil struct {       //将小学生的结构体与学生绑定
	stu student.Student              //直接在结构体内部建立嵌入匿名Student

}
func (p *Pupil) Testing(){
	fmt.Printf("小学生%v 正在考试\n",p.stu.Name)   //特有的方法字段保留	
}

func main(){
	var s student.Student = student.Student{
		Name:"小明",
		Age : 18,
	}
	s.Testing()
	s.SetScore(98)
	s.ShowInfo()
	
	p := &Pupil {} 
	p.stu.Name="小小明"  //调用父类的字段
	p.stu.Age = 10
	p.Testing()            //调用子类的testing方法
	p.stu.Testing()       //调用父类的testing方法
	p.stu.SetScore(98)
	p.stu.ShowInfo()
}
