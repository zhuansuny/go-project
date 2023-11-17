package main
import(
	"fmt"
	//"unsafe"

)

func test() bool{
	fmt.Println("test...")
	return true
}
func main(){	
	var a int = 97
	// var b float64 = 123
	// var c float64  
	// c = 5.0/9*(b-100)
	// fmt.Printf("还有%d个星期零%d天\n",a/7,a%7)
	// fmt.Printf("华氏温度%f对应摄氏度%f\n",b,c)
	// fmt.Printf("a的内容 %d 类型 %T 字节 %d\n",a,a,unsafe.Sizeof(a))
	if a<10&&test(){
		fmt.Println("ok.....")
	}else{
		fmt.Println("fail....")
	}

}