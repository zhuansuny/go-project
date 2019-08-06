package main //网络编程服务器

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)

		//1.等待客户端通过conn发送信息
		//2.如果客户端没有writer[发送]，那么协程会一直阻塞在这里
		fmt.Println("服务器在等待客户端发送信息", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端退出read err=", err)
			return
		}
		//3.显示客户端发送的内容到服务器的终端
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听。。。。")
	//tcp代表使用的网络协议是TCP
	//0.0.0.0:8888代表在本地监听8888端口
	//cmd netstat -ano 可以查看占用的端口号
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("服务器监听失败")
		return
	}
	defer listen.Close() //延时关闭listen
	//循环等待客户端连接
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept() err=", err)
		} else {
			fmt.Println("accept() success=", conn)
		}
		go process(conn)
	}

}
