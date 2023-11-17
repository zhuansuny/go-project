package main //数据库连接池
//事先初始化一定数量的链接，放入到链接池
//当go需要操作redis时，直接从链接池取出链接
//好处是节省临时获取Redis链接的时间，从而提高效率

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

func init() { //init函数在main函数之前执行
	pool = &redis.Pool{
		MaxIdle:     8,   //链接池最大空闲链接数
		MaxActive:   0,   //最大链接数，0表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接，链接哪个服务器
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func main() {
	conn := pool.Get() //从pool中取出一个链接
	defer conn.Close() //延时放回pool
	_, err := conn.Do("set", "name", "tom你好")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	//通过go向redis读取数据
	r, err := redis.String(conn.Do("get", "name")) //返回值的类型是空接口,需要redis.String转换为字符串
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	//应该
	fmt.Println("读取的数据为", r)
	fmt.Println("操作成功")

	pool.Close() //关闭链接池，关闭后无法再取出链接

}
