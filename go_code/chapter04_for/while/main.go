package main
import (
	"fmt"
)
//golang 没有while 需要使用for来实现
func main(){

// while的for实现
	i:=1
	for{
		if i>100 {
			break
		}
		fmt.Print("第",i)
		i++

	}
// do while的for实现
	i = 1
	for{
		fmt.Print("第",i)
		i++
		if i>100 {
			break
		}
	}


}