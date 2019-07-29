package  student
import (
	"fmt"
)

type Student struct {   
	Name string
	Age int
	Score int
}

func (p *Student) ShowInfo(){
	fmt.Printf("学生%v 年龄=%v 成绩=%v\n",p.Name,p.Age,p.Score)
}

func (p *Student) SetScore(score int ){
	p.Score = score
	
}
func (p *Student) Testing(){
	fmt.Printf("学生%v 正在考试\n",p.Name)	 
}
