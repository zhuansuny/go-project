package main  //map排序，golang中没有专门的函数来给map排序
import (
	"fmt"
	"sort"
)

func main(){
	map1 := make(map[int]int,10)
	map1[10] = 10
	map1[1] = 12
	map1[2] = 88
	map1[3] = 99



	fmt.Println(map1)

	var keys []int     //先定义一个切片，吧key遍历放到切片中，排序，再按照顺序取出
	for k, _ := range map1 {
		keys =append(keys,k)
		
	}
	sort.Ints(keys)      //排序
	fmt.Println(keys)

	for _, k := range keys {       //按照keys升序取出数据
		fmt.Printf("map[%d] = %d  ",k,map1[k])
	}
}