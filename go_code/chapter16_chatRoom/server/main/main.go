package main

//监听，等待客户端连接，以及初始化的工作
import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//需要延时关闭conn
	defer conn.Close()
	process := &Processor{
		Conn: conn,
	}
	err := process.process2()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	fmt.Println("服务器在8889端口监听。。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")

	defer listen.Close()
	if err != err {
		fmt.Println("监听失败，err=", err)
		return
	}

	for {
		fmt.Println("等待客户端来链接服务器。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接受失败，err=", err)
		}

		//一旦链接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}

}
