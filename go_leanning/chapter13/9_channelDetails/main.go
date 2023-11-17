package main //协程与管道结合2

import (
	"fmt"
	"time"
)

func test() { //如果起了一个协程，但是这个协程出现了panic，如果没有捕获这个panic，就会造成程序的崩溃
	defer func() {
		if err := recover(); err != nil { //可以使用recover来捕获panic进行处理，这样即使发生错误，但主线程不受影响
			fmt.Println("test发生错误", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang" //这里没有分配空间会报错

}

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello,world")
	}
}

func main() {

	//-------------------------只读与只写------------------------------------

	var chan1 chan int //默认是双向管道
	chan1 = make(chan int, 3)
	chan1 <- 10
	<-chan1              //可读可写
	var chan2 chan<- int //声明为只写
	chan2 = make(chan int, 3)
	chan2 <- 10

	// var chan3 <-chan int //声明为只读
	// chan3 = make(chan int, 3)

	//可以应用于函数中，将管道传给函数时可以定义为只读或只写
	//func test(chan2 chan<- int,chan3 <-chan int ) chan2和chan3都是chan int类型，但在函数中只能进行只读或只写

	//---------------------------select------------------------------------
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	//传统方法在遍历时，如果不关闭管道会阻塞而导致deadlock
	//在实际开发中，可能不好确定什么时候关闭管道
	//使用select，可以解决这个问题
loop:
	for {
		select {
		case v := <-intChan: //如果取不到数据，不会堵塞，会到下一个case匹配
			fmt.Printf("从intchan中读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringchan中读取的数据%d\n", v)
		default: //都取不到可以加入自己的业务逻辑
			fmt.Println("读取不到数据")
			break loop
		}
	}

	go test() //使用recover来捕获panic进行处理,不会影响sayHello协程执行
	go sayHello()
	time.Sleep(time.Second)
}
