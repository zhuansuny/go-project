package main
import(
	"fmt"
	"math/rand"
)

func main(){
	//循环随机生成整数，当生成99时使用break退出循环
	var num int
	var i int64
	for{
		rand.Seed(i)
		num=rand.Intn(100)+1
		fmt.Println("随机生成的数是",num)
		if num==99{
			break
		}
		i++	
	}
	fmt.Println(num)

	//break 默认跳出最近的for循环
	//break 后面可以指定标签，跳出标签对应的for循环 
	count := 3
	lable2:       //标签
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			for k := 0; k < count; k++ {
				if (k>1){
					break lable2  //指定标签
				}
				fmt.Print(i," ",j," ",k," \n")
			}
		}
	}

	
}