package main //全局变量加锁改进协程同步问题

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10) //定义一个全局map
	lock  sync.Mutex              //第一种全局变量加锁改进协程同步问题（低水平）  互斥锁
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()    //加锁
	myMap[n] = res //给map赋值
	lock.Unlock()  //解锁
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	time.Sleep(10 * time.Second) //需要主线程估算休眠时间
	// for i := 1; i < len(myMap); i++ {
	// 	fmt.Printf("mymap[%v]=%v\n", i, myMap[i])
	// }
	for i, v := range myMap {
		fmt.Printf("mymap[%v]=%v\n", i, v)
	}
}
