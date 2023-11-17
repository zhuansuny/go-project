package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 { //os.Args代表在CMD运行是在 .exe文件后输入的内容
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0]) //未输入内容，报错输出
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name) //判断输入的IP格式是否正确
	if addr == nil {
		fmt.Println("Invaild address")
	} else {
		fmt.Println("The address is", addr.String())
	}
	os.Exit(0)
}
