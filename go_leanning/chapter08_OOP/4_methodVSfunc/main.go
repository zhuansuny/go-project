package main  
import (
	"fmt"
)

type Person struct{
	name string
	age int
}

func test01(p Person){
	p.name ="garren"
	fmt.Println(p.name)
}
func test02(p *Person){
	p.name ="tom"
	fmt.Println(p.name)
}

func (p Person)test03(){
	p.name ="mary"
	fmt.Println(p.name)
}
func (p *Person) test04(){
	p.name ="teemo"
	fmt.Println(p.name)
}
func main(){
	var p Person
	p.name ="jack"
	test01(p)  //函数只能传入Person类型，不可传入指针test01(&p)
	test02(&p)  //函数只能传入*Person指针类型，不可传入test01(p)  可以改变p的值
	
	p.test03()
	(&p).test03()  //实例以及实例的指针都可调用test03  但都是值传递

	p.test04()
	(&p).test04()  //实例以及实例的指针都可调用test04  但都是引用传递 （主要取决于函数的定义）







}