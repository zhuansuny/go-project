package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func main() {
	monster1 := Monster{} //实例一个monster

	conn, err := redis.Dial("tcp", "127.0.0.1:6379") //连接redis数据库
	if err != nil {
		fmt.Println("连接redis数据库失败")
		return
	}
	fmt.Println("连接redis数据库成功")
	defer conn.Close() //延时关闭

	for i := 1; i < 4; i++ { //循环三次输入monster
		monster := "monster" + string(i)
		fmt.Printf("请按输入monster%d的 name\n", i)
		fmt.Scan(&monster1.Name)
		fmt.Printf("请按输入monster%d的 age\n", i)
		fmt.Scan(&monster1.Age)
		fmt.Printf("请按输入monster%d的 skill\n", i)
		fmt.Scan(&monster1.Skill)

		_, err = conn.Do("hmset", monster, "name", monster1.Name, "age",
			monster1.Age, "skill", monster1.Skill) //将monster添加到数据库

		if err != nil {
			fmt.Println("set err=", err)
			return
		}

		//fmt.Println("操作成功")
	}
	fmt.Println("遍历数据库结果为")
	for i := 1; i < 4; i++ {
		monster := "monster" + string(i)
		r, err := redis.Strings(conn.Do("hgetall", monster))
		if err != nil {
			fmt.Println("hmget err=", err)
			return
		}

		fmt.Printf("monster%d=%v\n", i, r)
	}
}
