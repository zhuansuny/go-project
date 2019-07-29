package main
import "fmt"
import "unsafe"
import "go_code/chapteor03/demo3/model"

func main(){
	var a int ='中'
	var b string = "中"
	var n = 100
	fmt.Printf("n1 的类型 %T" ,n)
	fmt.Println(b,a)
	var n2 int64 = 10
	fmt.Printf("n2的类型 %T n2占用的字节数是 %d",n2,unsafe.Sizeof(n2))
	fmt.Println(model.HeroName)
}