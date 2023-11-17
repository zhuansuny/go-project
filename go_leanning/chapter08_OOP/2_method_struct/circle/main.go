package main  //方法，通常用于结构体
import (
	"fmt"
)

type Circle struct {
	Radius float64       //字段使用大写，在别的包可以调用

}

func (c Circle) area()float64{   //该方法是结构体的值传递，方法内修改实例字段值，不会影响实例
	a :=3.14*c.Radius*c.Radius  //计算面积
	return a
}

func (c *Circle) area1()float64{ //加*代表引用传递，方法内修改实例字段值会影响实例
	a :=3.14*c.Radius*c.Radius  //计算面积 ,并且glang底层对此有优化可以用指针变量直接调用字段
	c.Radius = 5            
	return a
}

func main(){
	var c Circle
	c.Radius = 4     //给实例的字段半径赋值
	res := c.area()
	res1 := c.area1()  //底层优化可以不加&符号 res1 := &c.area1() 
	fmt.Println(res,res1,c.Radius)  //方法内修改实例字段值会影响实例,c.Radius输出为5

}