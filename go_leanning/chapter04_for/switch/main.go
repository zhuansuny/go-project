package main
import(
	"fmt"
)

func main(){
	var month byte 
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)
	//fmt.Printf("输入的字符是%c",char)
	switch month {
	case 3,4,5:
		fmt.Printf("%d月份是春季",month)
		fallthrough   //golang 中默认switch自带break，不需要手动添加，若想继续执行可加上fallthrough
	case 6,7,8:
		fmt.Printf("%d月份是夏季",month)
	case 9,10,11:
		fmt.Printf("%d月份是秋季",month)
	case 12,1,2:
		fmt.Printf("%d月份是冬季",month)
	default :
		fmt.Println("输入月份错误")
	}
	

}