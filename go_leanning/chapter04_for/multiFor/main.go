package main
import (
	"fmt"
)

func main(){
	var count int = 10
	//九九乘法表
	for i := 1; i < count; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d * %d = %d ",i,j,i*j)
		}
		fmt.Println()
	}
	//打印金字塔
	for i := 0; i < count; i++ {
		for k := 0; k < count-i-1; k++ {
			fmt.Print(" ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}