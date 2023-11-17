package main  //切片的动态追加
import (
	"fmt"
)
func main(){
	var arr [10]int =[...]int {2,3,3,1,4,5,3,8,9,1}
	slice := arr[0:3]
	fmt.Println("slice=",slice)
	slice2 :=append(slice,400,500,600)  //append的函数可以将slice追加元素，若它有足够的容量，其目标就会重新切片以容纳新的元素。否则，就会分配一个新的基本数组
	fmt.Println("slice2=",slice2)
	slice3 :=append(slice,slice2...)
	fmt.Println("slice3=",slice3)
	var slice4 =make([]int,10)
	copy(slice4,slice3)
	fmt.Println("slice4=",slice4)



}