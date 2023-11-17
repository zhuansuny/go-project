package main
import ("fmt")

func fbn(n int) ([]uint64){       //斐波那契数计算 1 ,1 ,3 ,5 ,8 ,13 当前数是前面两个数之和
	fbnSlice := make([]uint64,n)  //将斐波那契数放入切片中
	fbnSlice[0] =1
	fbnSlice[1] =1
	for i := 2; i < n; i++ {
		fbnSlice[i] =fbnSlice[i-1]+fbnSlice[i-2]
	}

	return fbnSlice
}

func middle(a *[10]int,l int,r int,n int){
	m := (l+r)/2
	if(l>r){
		fmt.Println("未找到")
		return
	}
	if(n==(*a)[m]){
		fmt.Println("找到下标为",m)
		
	
	}else if(n>(*a)[m]){
		middle(a,m+1,r,n)
		
	}else if(n<(*a)[m]){
		middle(a,l,m-1,n)
		
	}
	
}

// func BinaryFind(arr *[10]int, leftIndex int, rightIndex int, findVal int) {

// 	//判断leftIndex 是否大于 rightIndex
// 	if leftIndex > rightIndex {
// 		fmt.Println("找不到")
// 		return
// 	}

// 	//先找到 中间的下标
// 	middle := (leftIndex + rightIndex) / 2

// 	if (*arr)[middle] > findVal {
// 		//说明我们要查找的数，应该在  leftIndex --- middel-1
// 		BinaryFind(arr, leftIndex, middle - 1, findVal)
// 	} else if (*arr)[middle] < findVal {
// 		//说明我们要查找的数，应该在  middel+1 --- rightIndex
// 		BinaryFind(arr, middle + 1, rightIndex, findVal)
// 	} else {
// 		//找到了
// 		fmt.Printf("找到了，下标为%v \n", middle)
// 	}
// }
func main(){
	fubSlice :=fbn(10)
	fmt.Println(fubSlice)
	arr := [...]int{0,1,2,3,4,5,6,7,8,100}
	middle(&arr,0,len(arr) - 1, -8)
	
}