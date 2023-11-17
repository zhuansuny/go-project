package main
import(
	"fmt"
	"unsafe"
	"strconv"
)
func main(){

	var str string ="true"
	var b bool
	b,_=strconv.ParseBool(str)
	fmt.Printf("b type %T 内容 %v 字节%d\n",b,b,unsafe.Sizeof(b))
	str="1236155146"
	var num int64
	num,_ =strconv.ParseInt(str,10,64)
	fmt.Printf("num type %T 内容 %v 字节%d\n",num,num,unsafe.Sizeof(num))
	str="1.2"
	var num2 float64
	num2,_=strconv.ParseFloat(str,64)
	fmt.Printf("num2 type %T 内容 %v 字节%d\n",num2,num2,unsafe.Sizeof(num2))



}