package main //普通一个主线程进行计算对比

import (
	"fmt"
	"time"
)

func main() {
	primeChan := make(chan int, 20000)
	start := time.Now().Unix()
	for num := 1; num <= 100000; num++ {

		flag := true //假设是素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			//将这个数就放入到primeChan
			primeChan <- num
		}

	}
	end := time.Now().Unix()
	fmt.Println("普通的方法耗时=", end-start)

}
