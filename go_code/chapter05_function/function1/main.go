package main
import (
	"fmt"
)
func getSum(n1 int,n2 int) (int,int) {      //golang中函数可以有多个返回值在形参后，可以用_符号占位忽略
	sum := n1+n2                            //sum是函数的局部变量，只能在getsum中使用
	sum2 :=n1+n2+n2                         //基本类型以及数组默认是值传递（java中数组是引用类型）
	return sum,sum2                         //若想修改函数外数据，可以将变量地址指针&传入
}                                          //golang 不支持传统的函数重载（即通过形参的个数及类型区分）



func getSum2(n1 int,n2 int) (sum int,sub int) {    //可以为返回值取名
	sum = n1+n2                                 //默认定义了sum,sub 函数内不用再定义可以直接使用
	sub =n1+n2+n2                        
	return                                      //可以直接返回
}

func myFun(funvar func(int, int) (int,int),num1 int ,num2 int ) (int,int) {                       
	return  funvar(num1,num2)                      //函数可以当作形参传入函数中
}   

func sum1(args... int)int{          //可以定义不定个数形参的函数(可变函数)
	var sum1 int
	for i := 0; i < len(args); i++ { //取出啊args的个数来确定输入的个数
		sum1 = sum1+args[i]            //求出输入的值之和
	}
	return sum1
}


func main(){
	sum,sum2 := getSum(10,20)
	sum3,_ :=getSum(34,20)        //可以使用_符号占位忽略
	fmt.Println("sum=",sum)
	fmt.Println("sum2=",sum2)
	fmt.Println("sum3=",sum3)

	a := getSum                  //函数可以当作一个引用类型，将a声明为getSum
	res,_ := a(10,40)            //可以通过a调用getSum函数
	fmt.Println("res=",res)

	res2,_ := myFun(a,11,22)    //函数可以当作形参传入函数中
	fmt.Println("res2=",res2)
 
	type myInt int        //自定义数据类型，虽然两个都是int类型，但golang认为两个是不同类型

	var num1 myInt
	var num2 int
	num1 = 40
	num2 =int(num1)    //默认不同类型，需要强制转换
	fmt.Println(num1,num2)
	num3 :=sum1(1,2,3,8)
	fmt.Println("num3=",num3)
}