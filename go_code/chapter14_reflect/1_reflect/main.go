package main //反射

import (
	"fmt"
	"reflect"
)

type Stu struct {
	name string
	age  int
}

func reflectTest01(b interface{}) {
	//通过反射获取传入的变量type，kind，值
	//1.先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp =", rTyp)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b) //值是10，但不是int的类型，是reflect.ValueOf类型
	fmt.Println("rVal =", rVal)
	num := 2 + rVal.Int() //可以用rVal.Int()方法转成int类型
	fmt.Println("num =", num)
	//------------再转换会int类型---------------
	iV := rVal.Interface() //将rVal转换为空接口类型
	num2 := iV.(int)       //类型断言转换为int
	fmt.Println("num2 =", num2)

}

func reflectTest02(b interface{}) {
	//通过反射获取传入的变量type，kind，值
	//1.先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp =", rTyp)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal =", rVal)

	//------------再转换回int类型---------------
	iV := rVal.Interface() //将rVal转换为空接口类型
	fmt.Printf("iv =%v 类型是%T\n", iV, iV)
	stu, ok := iV.(Stu) //类型断言转换为stu
	if ok {
		fmt.Printf("stu =%v 类型是%T\n", stu.name, stu)
	}
	kind1 := rVal.Kind()
	kind2 := rTyp.Kind()
	fmt.Printf("kind =%v,kind =%v\n", kind1, kind2) //类型都是struct
}

func reflectTest03(b interface{}) {

	rVal := reflect.ValueOf(b) //值是15
	rVal.Elem().SetInt(20)
	fmt.Println("rVal =", rVal)

}

func main() {
	//反射int
	var num int = 10
	reflectTest01(num)
	//反射Stu结构体
	stu := Stu{
		name: "jack",
		age:  18,
	}
	reflectTest02(stu)

	num2 := 15 //通过传入地址修改num2的值
	reflectTest03(&num2)
	fmt.Println(num2)

}
