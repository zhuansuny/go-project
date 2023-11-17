package main //redis数据库hash指令

import (
	"fmt"

	"github.com/garyburd/redigo/redis" //引入redis包
	//需要在gopath目录下cmd 运行 go get github.com/garyburd/redigo/redis 安装go的redis包
)

func main() {
	//连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接redis数据库失败")
		return
	}
	fmt.Println("连接redis数据库成功")
	defer conn.Close()

	//通过go向redis写入数据
	_, err = conn.Do("hmset", "user", "name", "teemo", "age", 18, "id", 1)
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	//通过go向redis读取数据
	r, err := redis.Strings(conn.Do("hmget", "user", "name", "age", "id")) //返回值的类型是空接口,需要redis.String转换为字符串
	if err != nil {
		fmt.Println("hmget err=", err)
		return
	}
	//应该
	fmt.Println("读取的数据为", r)
	fmt.Println("操作成功")
}
