package main  //二维数组
import (
	"fmt"
)
func main(){
	var arr [2][3]int                  //二维数组的定义
	var arr2 = [2][3]int{{1,2,3},{4,5,6}}
	var arr3 = [...][]int{{1,2,3},{4,5,6}}
	arr4 := [][]int{{1,2,3},{4,5,6}}
	arr[1][1] = 10
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Printf("arr地址是%p\n",&arr)        //arr 的地址与arr[0]和arr[0][0]相同
	fmt.Printf("arr[0]地址是%p\n",&arr[0])
	fmt.Printf("arr[0]地址是%p\n",&arr[0][0])
	fmt.Printf("arr[1]地址是%p\n",&arr[1]) //arr[1]的地址比arr[0]大12个字节
	fmt.Printf("arr[0]的长度是%d\n",len(arr[0])) //arr[]的长度为行数3
	fmt.Printf("arr的长度是%d\n",len(arr))  //arr的长度为列数2
	//

	//二维数组的遍历
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("arr[%d][%d] = %d  ",i,j,arr[i][j]) //for
		}
		fmt.Println()
	}

	for i, v := range arr {
		for j, v1 := range v {
			fmt.Printf("arr[%d][%d] = %d  ",i,j,v1)  //for-range
		}
		fmt.Println()
	}

}