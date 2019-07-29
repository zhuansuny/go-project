package main
import(
	"fmt"
	"unsafe"
) 
func main(){
	var c1 byte =100
	var c2 int16 =30000
	var b bool = 1
	fmt.Println(b)
	fmt.Printf("n2的类型 %T n2占用的字节数是 %d \n",b,unsafe.Sizeof(b))
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Printf("c1=%c",c2)
	fmt.Println()
	fmt.Printf("n2的类型 %T n2占用的字节数是 %d",c2,unsafe.Sizeof(c2))
	

}