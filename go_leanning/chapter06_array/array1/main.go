package main
//一维数组在golang中是值类型
//数组定义后长度和类型是固定的，不能变化
//数组中元素可以是值类型和引用类型，但不能混用
//golang中数组是值传递，传入函数不会影响原数组
import (
	"fmt"
)

func main (){
	var hens [6]float64   //一维数组定义
	hens[0] = 3.0         //一维数组赋值
	hens[1] = 5.0
	hens[2] = 8.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	fmt.Println(hens)   //数组是值类型
	fmt.Printf("%p\n",&hens)   
	fmt.Println(&hens[0])   //数组的地址就是第一个元素的地址
	fmt.Println(&hens[2])   //数组的第二个元素地址就是第一个元素的地址加8
	total := 0.0
	for i := 0; i < len(hens); i++ {
		total += hens[i]    //一维数组遍历
	}

	for i ,value := range hens {  //一维数组遍历2   i是下标，若不需要可用_下划线忽略
		fmt.Printf("值是%v 下标是%v\n",value,i)
	}
	avg := fmt.Sprintf("%.2f",total/float64(len(hens)))  //将结果四舍五入保留2位小数
	fmt.Println(avg)



	//四种初始化数组的方式

	var numArr1 [3]int =[3]int{1,2,3}
	fmt.Println("numArr1=",numArr1)     
	
	var numArr2 =[3]int{4,5,6}
	fmt.Println("numArr2=",numArr2)

	var numArr3 =[...]int{1,2,3}
	fmt.Println("numArr3=",numArr3)

	var numArr4 =[...]int{1:100,0: 2002,2:128}  //指定下标
	fmt.Println("numArr4=",numArr4)
}