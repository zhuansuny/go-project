package main 
			 
import (
	"fmt"
)

type Student struct{
	Name string
}

func TypeJudge (items... interface{}){       //断言最佳实践 判断传入参数的类型
	for index, v := range items {
		switch v.(type) {
		case bool :
			fmt.Printf("第%v个参数是 bool类型，值是%v\n",index+1,v)
		case float32 :
			fmt.Printf("第%v个参数是 float32 类型，值是%v\n",index+1,v)
		case float64 :
			fmt.Printf("第%v个参数是 float64 类型，值是%v\n",index+1,v)
		case int,int32,int64 :
			fmt.Printf("第%v个参数是 整数类型，值是%v\n",index+1,v)
		case string :
			fmt.Printf("第%v个参数是 string类型，值是%v\n",index+1,v)
		case Student :                                               //也可以断言自定义类型
			fmt.Printf("第%v个参数是 Student类型，值是%v\n",index+1,v)
		case *Student :
			fmt.Printf("第%v个参数是 *Student类型，值是%v\n",index+1,v)
		default :
			fmt.Printf("第%v个参数是 不确定类型，值是%v\n",index+1,v)

			
		}
		
	}
}


func main(){

	var n1 float32 =1.1
	var n2 float64 =2.1
	var n3 int = 3
	var name string ="teemo"
	var stu =Student{"garren"}
	var stu1 =&Student{"jack"}
	TypeJudge(n1,n2,n3,name,stu,stu1)


}