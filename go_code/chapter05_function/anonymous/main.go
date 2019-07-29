package main 
import (
	"fmt"
)
//匿名函数

var (
	fun1 =func(n1 int,n2 int)int{    //全局匿名函数
		return n1*n2
	}
)
func main(){
	res :=func(n1 int,n2 int)int{  //1.定义匿名函数的时候就直接调用，这种方式只能调用一次
			return n1+n2		   //该函数求两者之和
		}(10,20)

		fmt.Println("res=",res)

	a := func(n1 int,n2 int)int{  //1.定义匿名函数的时候赋值给变量a，这种方式可以反复调用
			return n1-n2		   //该函数求两者之和
		}
		
		res2 := a(20,10)
		res3 := a(30,10)
		fmt.Println("res2=",res2)
		fmt.Println("res3=",res3)
		res4 := fun1(10,20)      //全局匿名函数调用
		fmt.Println("res4=",res4)  
}