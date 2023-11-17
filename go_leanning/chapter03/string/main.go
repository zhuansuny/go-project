package main
import(
	"fmt"
	"unsafe"
	"strconv"
) 

var age int = 13  //全局变量这种方式正常
//name :="tom"      //这种方式等于两条语句 var name string; name="tom";在函数外不可赋值，这种方式错误

func main(){
	var num int32 =888
	var num2 float64 = 12.3654
	var b bool = true
	var mychar byte ='h'
	var str string
	str = fmt.Sprintf("%d",num )
	fmt.Printf("str的类型 %T 占用的字节数 %d 内容是%s\n",str,unsafe.Sizeof(str),str)
	str = fmt.Sprintf("%f",num2 )
	fmt.Printf("str的类型 %T 占用的字节数 %d 内容是%s\n",str,unsafe.Sizeof(str),str)
	str = fmt.Sprintf("%t",b )
	fmt.Printf("str的类型 %T 占用的字节数 %d 内容是%s\n",str,unsafe.Sizeof(str),str)
	str = fmt.Sprintf("%c",mychar )
	fmt.Printf("str的类型 %T 占用的字节数 %d 内容是%s\n",str,unsafe.Sizeof(str),str)
  
	var num3 int = 99
	var num4 float64 =23.456
	var b2 bool = true

	str =strconv.FormatInt(int64(num3),10)
	fmt.Printf("类型 %T 字节数 %d 内容%s\n",str,unsafe.Sizeof(str),str)
	str = strconv.FormatFloat(num4,'f',8,64)
	fmt.Printf("类型 %T 字节数 %d 内容%s\n",str,unsafe.Sizeof(str),str)
	str = strconv.FormatBool(b2)
	fmt.Printf("类型 %T 字节数 %d 内容%s\n",str,unsafe.Sizeof(str),str)
}