package main  //数组测试
import (
	"fmt"
)

/*func insertArray(arr *[10]int, n int)[11]int{  //在一个有序数组中插入一个数据，并返回一个有序数组
	//l :=len(*arr)+1
	var arr1 [11]int                             //定义一个数组存放数据
	if(*arr)[len(*arr)-1] <= n{                  //判断是否大于最大值
		for i := 0; i < len(arr1)-1; i++ {   
			arr1[i] = (*arr)[i]
		}
		arr1[len(arr1)-1] = n
		return arr1
	} 
	if(*arr)[0] >= n{                        //判断是否小于最小值
		for i := 0; i < len(arr1)-1; i++ {
			arr1[i+1] = (*arr)[i]
		}
		arr1[0]=n
		return arr1
	}
	for i := 1; i < len(*arr)-1; i++ {  
		if (*arr)[i] >= n{             //找到大于插入数据的值的下标
			for j := 0; j < i; j++ {   //将下标之前的数值赋给新数组
				arr1[j] = (*arr)[j]   
			}
			arr1[i] = n                    //将n赋给下标数组
			for j := i; j < len(arr1)-1; j++ {///下标之后的数组加一赋给新数组
				arr1[j+1] = (*arr)[j]
			}
			return arr1             //找到就返回
		}
	}
	return arr1

} */

func insertArray(arr *[]int, n int)[]int{  //在一个有序数组中插入一个数据，并返回一个有序数组
	//l :=len(*arr)+1
	var arr1 []int  =  append((*arr),n)                       //定义一个数组存放数据
	if(*arr)[len(*arr)-1] <= n{                  //判断是否大于最大值
		arr1 = append((*arr),n)
	} 
	if(*arr)[0] >= n{                        //判断是否小于最小值
		for i := 0; i < len(arr1)-1; i++ {
			arr1[i+1] = (*arr)[i]
		}
		arr1[0]=n
		return arr1
	}
	for i := 1; i < len(*arr)-1; i++ {  
		if (*arr)[i] >= n{             //找到大于插入数据的值的下标
			for j := 0; j < i; j++ {   //将下标之前的数值赋给新数组
				arr1[j] = (*arr)[j]   
			}
			arr1[i] = n                    //将n赋给下标数组
			for j := i; j < len(arr1)-1; j++ {///下标之后的数组加一赋给新数组
				arr1[j+1] = (*arr)[j]
			}
			return arr1             //找到就返回
		}
	}
	return arr1

}



func main(){
	var arr =[]int{2,3,6,8,12,23,43,56,89,98}
	arr1 :=insertArray(&arr,88)
	fmt.Println(arr1)

}