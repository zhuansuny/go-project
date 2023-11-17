package main //管道本质是一个数据结构-队列
//数据是先进先出fifo
//管道是线程安全的，不需要加锁
//管道是有类型的，一个string的管道只能存放string类型的数据
import (
	"fmt"
)

func main() {
	//-----------------定义一个int的管道----------------------
	var intChan chan int
	intChan = make(chan int, 3) //make可以存放3个int的管道

	fmt.Printf("intChan的值是%v，intChan本身的地址是%p\n", intChan, &intChan)

	intChan <- 10 //向管道写入数据
	num := 211
	intChan <- num //写入数据时不可超出容量
	intChan <- 150
	close(intChan) //将int管道关闭，不能再添加数据，但可以取出 ，取出后也不可以再添加
	fmt.Printf("intChan  len是%v，cap是%v\n", len(intChan), cap(intChan))

	//channel的遍历
	for v := range intChan { //若管道没有关闭遍历会报错
		fmt.Println("v=", v)
	}

	//从管道取出数据
	num2, ok := <-intChan // 管道关闭后没有数据时再取数据为0,ok为false，管道未关闭则会报错

	fmt.Println(num2, ok)

	//  ---------------------定义一个map的管道---------------
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)

	var m = make(map[string]string)
	m["name"] = "牛魔王"
	mapChan <- m
	m1 := <-mapChan
	fmt.Println(m1["name"])

	//---------------------定义一个空接口管道----------------

	allChan := make(chan interface{}, 10)
	allChan <- m
	allChan <- 100

	m2 := <-allChan
	//fmt.Println(m2["name"]) //取出默认是空接口，不可调用map的方法，可以使用类型断言
	a := m2.(map[string]string) //类型断言
	fmt.Println(a["name"])

}
