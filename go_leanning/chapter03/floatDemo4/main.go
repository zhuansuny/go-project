package main
import(
	"fmt"
	"unsafe"
) 
func main(){
	var n2 float64 = 10001.388
	fmt.Println(n2)
	fmt.Printf("n2的类型 %T n2占用的字节数是 %d",n2,unsafe.Sizeof(n2))

}