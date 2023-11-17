package main  //反序列化

import (
	"fmt"
	"encoding/json"
)

type Monster struct{
	Name string `json:"name"`  //`json:"name"` tag标签代表在json序列化是转成name(大写在别的程序中可能不兼容)
	Age int `json:"age"`
	Birthday string     //不加tag标签，将会显示原本的Birthday
	Skill string
}

//-------------------------Json的反序列化-结构体-------------------------

func UnmarshalStruct(){
	 //需要对双引号进行转义处理 ，实际运行程序获取的Json字符串不需要转义
	str :="{\"name\":\"牛魔王\",\"age\":500,\"Birthday\":\"0321\",\"Skill\":\"牛魔拳\"}"  

	var monster Monster
	//反序列化
	err :=json.Unmarshal([]byte(str),&monster)
	if err !=nil{
		fmt.Printf("反序列化失败，err=%v\n",err)
	}
	fmt.Printf("反序列化后，moster=%v\n moster.Name=%v\n",monster,monster.Name)
} 

//-------------------------Json的反序列化-Map-------------------------
func UnmarshalMap(){
	str :="{\"name\":\"白骨精\",\"age\":500,\"Birthday\":\"0321\",\"Skill\":\"牛魔拳\"}"

	var a map[string]interface{}  //反序列化不需要make,因为make封装到 Unmarshal函数中
	//反序列化
	err :=json.Unmarshal([]byte(str),&a)
	if err !=nil{
		fmt.Printf("反序列化失败，err=%v\n",err)
	}
	fmt.Printf("反序列化后，a=%v\n a[name]=%v\n",a,a["name"])
} 

//-------------------------Json的反序列化-切片-------------------------


func UnmarshalSlice(){
	str :="[{\"address\":\"白骨山\",\"age\":300,\"name\":\"白骨精\"},{\"address\":\"黑风山\","+
		"\"age\":400,\"name\":\"黑熊精\"}]"

	var slice []map[string]interface{}  //反序列化不需要make,因为make封装到 Unmarshal函数中
	//反序列化
	err :=json.Unmarshal([]byte(str),&slice)
	if err !=nil{
		fmt.Printf("反序列化失败，err=%v\n",err)
	}
	fmt.Printf("反序列化后，slice=%v\n",slice)
} 



func main(){

	//要保证反序列化之后的数据类型和序列化之前的一致
	UnmarshalStruct()
	UnmarshalMap()
	UnmarshalSlice()
	


}