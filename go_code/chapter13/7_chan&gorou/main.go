package main //协程与管道结合

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ { //编译器会分析，如果只有写没有读，超过容量就会，则会有阻塞报deadlock错误
		intChan <- i //管道写入数据  //如果有读有写，写管道和读管道的频率不一致，等待取出数据后再会写入
		fmt.Printf("writeData写的数据=%v\n", i)
		time.Sleep(time.Microsecond)
	}

	close(intChan) //关闭管道,当读完数据后ok标志为false退出
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan //对于同一通道，接收操作是阻塞的，直到发送者可用。如果通道中没有数据，接收者会保持阻塞。
		if !ok {
			break
		}
		fmt.Printf("readData读到的数据=%v\n", v)
		time.Sleep(time.Second)
	}
	exitChan <- true
	close(exitChan)

}

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)

	//time.Sleep(time.Second) //传统方法，主线程休眠等待协程执行完毕
	for {
		_, ok := <-exitChan //管道方法，不停去读取
		if !ok {
			break
		}
	}
}
