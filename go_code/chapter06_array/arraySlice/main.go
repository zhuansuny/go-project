package main  //切片
import (
	"fmt"
)
func main(){
	var arr [5]int = [...]int{22,33,11,44,555}
	fmt.Println("arr=",arr)
	fmt.Printf("slice=%p\n",&arr[1])
	slice := arr[1:3]               //slice是引用，将arr的第2到第三个元素及地址赋给slice(不含3)
	//slice := arr[:3]           //start省略代表从0开始
	//slice := arr[1:]           //end省略代表从start取完
	//slice := arr[:]           //都省略代表全部取完
	fmt.Printf("slice=%p\n",&slice[0])  //slice的第一个元素的地址与arr数组的第二个元素的地址相同
	slice[1] = 10                  //修改slice会影响arr数组
	fmt.Println("arr=",arr)
	fmt.Println("slice的长度为",len(slice))
	fmt.Println("slice的容量为",cap(slice))


	//切片的另外两种方式
	var slice1 []int = make([]int,4,10)  //make([]类型，初始化0的个数，容量长度)
	slice1[0] =10
	slice1[1] =20
	slice2 := slice1[1:9]             //切片可以继续再进行切片
	fmt.Println("slice1的值为",slice1)
	fmt.Println("slice2的值为",slice2)
	fmt.Println("slice1的长度为",len(slice1))
	fmt.Println("slice1的容量为",cap(slice1))
	fmt.Println("slice2的长度为",len(slice2))
	fmt.Println("slice2的容量为",cap(slice2))

	
	//
	var strSlice []string = []string{"tom","jack","mary"}
	fmt.Println("strSlice的值为",strSlice)
	fmt.Println("strSlice的长度为",len(strSlice))
	fmt.Println("strSlice的容量为",cap(strSlice))

	//切片的遍历
	for i :=0 ;i<len(slice);i++ {                //普通for
		fmt.Printf("slice[%v] = %v\n",i,slice[i])	
	}

	for i, v := range slice{                  //for-range
		fmt.Printf("slice[%v] = %v\n",i,v)
	}

}