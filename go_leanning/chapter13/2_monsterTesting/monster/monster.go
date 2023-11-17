package model //单元测试实例

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct { //定义一个Monster结构体
	Name  string
	Age   int
	Skill string
}

func (monster *Monster) Store() bool { //定义Monster方法 序列化并保存到文件中
	data, err := json.Marshal(&monster) //序列化
	if err != nil {
		fmt.Printf("序列化失败，err=%v\n", err)
		return false
	}
	filePath := "d:/monster.txt"                                      //保存文件路径
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) //打开并创建文件
	if err != nil {
		fmt.Printf("打开文件失败，err=%v\n", err)
		return false
	}
	file.WriteString(string(data)) //将字符串写入文件
	return true

}

func (monster *Monster) ReStore() bool { //定义Monster方法 从文件读取字符串并反序列化到Monster的实例

	filePath := "d:/monster.txt"              //读取的文件路径
	content, err := ioutil.ReadFile(filePath) //读取文件
	if err != nil {
		fmt.Println("read file err=", err)
		return false
	}
	//var monster2 Monster
	json.Unmarshal(content, &monster) //将读取的字节反序列化到monster实例上

	return true

}

// func main() {
// 	monster := Monster{
// 		Name:  "牛魔王",
// 		Age:   199,
// 		Skill: "牛魔拳",
// 	}
// 	monster.Store()
// 	monster.Name = "狐狸精"
// 	fmt.Println(monster.Name)
// 	monster.ReStore()
// 	fmt.Println(monster.Name)

// }
