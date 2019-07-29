package main
import(
	"fmt"
)

func main(){

	//break 默认跳出最近的for循环
	//break 后面可以指定标签，跳出标签对应的for的当前循环
	count := 3
	lable2:       //标签
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			for k := 0; k < count; k++ {
				if (k>1){
					continue lable2  //指定标签
				}
				fmt.Print(i," ",j," ",k," \n")
			}
		}
	}

	for i := 0; i < 100; i++ {
		if i%2==0{
			continue
		}
		fmt.Println("奇数",i)
	}

	
}