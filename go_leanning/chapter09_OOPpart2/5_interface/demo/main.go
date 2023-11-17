package main 
//接口是对继承的补充
import (
	"fmt"
	"sort"
	"math/rand"
)


type Hero struct{   //定义一个结构体
	Name string
	Age int
}

type HeroSlice []Hero  //定义一个Hero切片类型

func (hs HeroSlice) Len() int{   //实现了自带data接口{  Len() Less() Swap()} 就可以使用sort方法
	return len(hs)               //求切片的厂长度
}

func (hs HeroSlice) Less(i,j int)bool {//判断切片两个下标的值得大小并返回布尔值
	return hs[i].Age<hs[j].Age
}


func (hs HeroSlice) Swap(i,j int){  //交换
	// hero := hs[i]
	// hs[i] = hs[j]
	// hs[j] = hero

	hs[i],hs[j] = hs[j],hs[i]  //交换可以简化写出一条语句

}

func main(){
	var slice = []int{21,33,11,3,1,0,-1,-8}
	sort.Ints(slice)
	fmt.Println(slice)

	var hs HeroSlice

	for i := 0; i < 10; i++ {
		//rand.Seed(int64 (i))
		hero := Hero{
			Name : fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age : rand.Intn(100) ,
		}
		hs = append(hs,hero)
	}
	fmt.Println("------------排序前---------------")
	for _, v := range hs {
		fmt.Println(v)
	}

	sort.Sort(hs)
	fmt.Println("------------排序后---------------")
	for _, v := range hs {
		fmt.Println(v)
	}
}