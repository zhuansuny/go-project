package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU() //获取本地逻辑cpu个数
	runtime.GOMAXPROCS(cpuNum) //设置多个cpu工作模式   go1.8以后不用设置

	fmt.Println("num=", cpuNum)

}
