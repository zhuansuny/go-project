package main //Json  https://www.json.cn 可以验证格式是否正确

/* [{"name":"小名","age":18, "address":["上海","北京"]},
 {"name":"小名","age":18, "address":["上海","北京"]}
 ] */   

 //上面是一个Json字符串的标准格式 是key-value键值对结构

import (
	"fmt"
	"encoding/json"
)

//先定义一个结构体

type Monster struct{
	Name string `json:"name"`  //`json:"name"` tag标签代表在json序列化是转成name(大写在别的程序中可能不兼容)
	Age int `json:"age"`
	Birthday string     //不加tag标签，将会显示原本的Birthday
	Skill string
}
//-------------------Json的序列化-结构体------------------
func testStruct(){
	//建立一个实例
	monster :=Monster{
		Name : "牛魔王",  
		Age :500 ,
		Birthday: "0321",
		Skill : "牛魔拳",
	}
	//Json的序列化
	data ,err :=json.Marshal(&monster)
	if err !=nil {
		fmt.Printf("序列化失败，err=%v\n",err)
		return
	}
	fmt.Printf("monster序列化后=%v\n",string(data))
}
//-------------------Json的序列化-map------------------
func testMap(){
	//定义一个map
	var a map[string]interface{}
	//使用map前先分配地址
	a =make(map[string]interface{})
	a["name"] ="红孩儿"
	a["age"] =30
	a["address"] = "洪崖洞"
	//Json的序列化
	data ,err :=json.Marshal(a)
	if err !=nil {
		fmt.Printf("序列化失败，err=%v\n",err)
		return
	}
	fmt.Printf("a map序列化后=%v\n",string(data))
}
//-------------------Json的序列化-切片------------------
func testSlice(){
	//定义一个map和切片
	var slice []map[string]interface{}
	var a map[string]interface{}
	//使用map前先分配地址
	a =make(map[string]interface{})
	a["name"] ="白骨精"
	a["age"] =300
	a["address"] = "白骨山"
	slice =append(slice,a)


	var a1 map[string]interface{}
	//使用map前先分配地址
	a1 =make(map[string]interface{})
	a1["name"] ="黑熊精"
	a1["age"] =400
	a1["address"] = "黑风山"
	slice =append(slice,a1)
	//Json的序列化
	data ,err :=json.Marshal(slice)
	if err !=nil {
		fmt.Printf("序列化失败，err=%v\n",err)
		return
	}
	fmt.Printf("slice序列化后=%v\n",string(data))
}

//-------------------Json的序列化-基本类型------------------

//基本类型的序列化就是转换成字符串，意义不大
func testFloat64(){
	var num1 float64 =23345.23

	data ,err :=json.Marshal(num1)
	if err !=nil {
		fmt.Printf("序列化失败，err=%v\n",err)
		return
	}
	fmt.Printf("float64序列化后=%v\n",string(data))


}

func main(){

//-------------------Json的序列化（以结构体、map、切片为例）------------------
	testStruct()
	testMap()
	testSlice()
	testFloat64()
	

}