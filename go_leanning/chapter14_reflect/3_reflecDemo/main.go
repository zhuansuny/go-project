package main //反射练习

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string
}

//方法，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//方法， 接收四个值，给结构体实例赋值
func (s Monster) Set(name string, age int, skill string) {
	s.Name = name
	s.Age = age
	s.Skill = skill
}

//方法，显示结构体实例的值
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end----")
}

func testStruct(a interface{}) {
	typ := reflect.TypeOf(a)  //获取到 reflect.Type
	val := reflect.ValueOf(a) //获取到reflect.Value
	kd := typ.Kind()          //获取a的类型
	if kd != reflect.Struct {
		fmt.Println("不是结构体")
		return
	}
	num := val.NumField() //获取该结构体有多少个字段
	fmt.Println("结构体的字段数目为", num)

	for i := 0; i < num; i++ {

		fmt.Printf("field %d : 值为= %v\n", i, val.Field(i))
		//获取到struct标签 `json:"name"` ,需要使用reflect.Type来获取tag标签
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("field %d: tag=%v\n", i, tagVal) //将标签值输出
		}
	}

	methodNum := val.NumMethod() //获取该结构体有多少个方法
	fmt.Println("结构体的方法数目为", methodNum)
	val.Method(1).Call(nil) //获取到第二个方法。调用它 //按照首字母排序

	//调用结构体的第1个方法Method(0)
	var params []reflect.Value //调用的传参需要放入到reflect.Value切片中
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())

}

func main() {
	monster := Monster{
		Name:  "狐狸精",
		Age:   18,
		Skill: "魅惑",
	}
	testStruct(monster)
}
