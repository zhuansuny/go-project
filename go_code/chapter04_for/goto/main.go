package main
import (
	"fmt"
)

//golang 支持goto语句，但不建议使用,容易造成程序混乱，调试困难
func main(){
	var n int =30
	fmt.Println("ok1")
	if n>20{
		goto label1 //满足条件跳转到16行
	}
	
	fmt.Println("ok2")
	fmt.Println("ok3")
	label1:
	fmt.Println("ok4")
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")

}