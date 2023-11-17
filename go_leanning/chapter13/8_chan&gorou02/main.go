package main //协程与管道结合2

import (
	"fmt"
	"time"
)

func putNum(intChan chan int, num int) {
	for i := 0; i <= num; i++ {
		intChan <- i

	}
	close(intChan) //关闭管道
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		//time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true //假设是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num

		}

	}

	fmt.Println("一个primeNum协程取不到数据，退出")
	exitChan <- true

}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 4)
	start := time.Now().Unix()
	go putNum(intChan, 100000)

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("运行时间是", end-start)
		close(primeChan)
	}()

	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Println("素数=", num)
	}
	fmt.Println("主线程退出")
}
