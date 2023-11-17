package main  //defer 是延时机制
import (
	"fmt"
)
//可以在defer后添加释放资源语句，以便于函数执行完毕后及时释放资源
func sum(n1 int,n2 int) int {  
	//当执行到defer时，暂不执行，会将defer后面的语句压入独立的栈中（defer栈）  
	//当函数执行完毕后，再从defer栈。按照先入后出的方式出栈执行
	defer fmt.Println("n1=",n1)   //并且将n1（=10）当前值入栈
	defer fmt.Println("n2=",n2)
	n1= n1+2                    //加2不会影响上面defer的n1值 
	sum := n1+n2 
	defer fmt.Println("sum=",sum)                                              
	return sum                      
}

func main(){
	
	fmt.Println("函数结果",sum(10,29))
}