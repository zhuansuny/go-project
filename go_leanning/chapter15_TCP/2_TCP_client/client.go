package main //网络编程客户端

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败")
		return
	}
	fmt.Println("connect success")
	//1.客户端发送单行数据
	reader := bufio.NewReader(os.Stdin) //os.Stdin代表标准输入 （终端）
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("读取数据失败", err)
		}

		if line == "exit\r\n" {

			return
		}
		fmt.Printf("读取的数据为%T", line)
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write失败", err)
		}
		fmt.Printf("客户端发送了%d个字节\n", n)

	}

}
