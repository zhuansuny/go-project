package main  //map切片实现动态增长
import (
	"fmt"
)

func main(){

	var monsters []map[string]string
	monsters = make([]map[string]string,2)
	if monsters[0]==nil{
		monsters[0] =make(map[string]string,2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "1500"
	}
	if monsters[1]==nil{
		monsters[1] =make(map[string]string,2)
		monsters[1]["name"] = "狐狸精"
		monsters[1]["age"] = "500"
	}

	 	// monsters[2] =make(map[string]string,2)    //不可以动态增加，会报错越界
		// monsters[2]["name"] = "狐狸精"
		// monsters[2]["age"] = "400"
	fmt.Println(monsters)

	//使用append函数动态增长

	newMonster :=map[string]string{
		"name" : "新的妖怪" ,
		"age"  : "200",
	}
	
	monsters = append(monsters,newMonster)  //把新妖怪加到map中
	fmt.Println(monsters)

}