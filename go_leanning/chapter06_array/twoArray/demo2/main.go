package main  //二维数组测试
import (
	"fmt"
)

func insertArray(arr *[10]int, n int)[11]int{  //在一个有序数组中插入一个数据，并返回一个有序数组
	//l :=len(*arr)+1
	var arr1 [11]int
	if(*arr)[len(*arr)-1] <= n{
		for i := 0; i < len(arr1)-1; i++ {
			arr1[i] = (*arr)[i]
		}
		arr1[len(arr1)-1]=n
		return arr1
	}
	if(*arr)[0] >= n{
		for i := 0; i < len(arr1)-1; i++ {
			arr1[i+1] = (*arr)[i]
		}
		arr1[0]=n
		return arr1
	}
	for i := 1; i < len(*arr)-1; i++ {
		if (*arr)[i] >= n{
			for j := 0; j < i; j++ {
				arr1[j] = (*arr)[j]
			}
			arr1[i] = n
			for j := i; j < len(arr1)-1; j++ {
				arr1[j+1] = (*arr)[j]
			}
			return arr1
		}
	}
	return arr1

}
func main(){
	var arr =[...]int{2,3,6,8,12,23,43,56,89,98}
	arr1 :=insertArray(&arr,88)
	//arr1 := &arr
	fmt.Println(arr1)

}