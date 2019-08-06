package main

import (
	"encoding/binary"
	"encoding/json"
	_ "errors"
	"fmt"
	"go_code/chapter16_chatRoom/commend/message"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Println("读取服务端发送的数据。。。")
	//conn.Read 只有在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭conn，就不会堵塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		if err == io.EOF {
			fmt.Println("客户端退出，服务器协程也退出", err)
			return
		}
		fmt.Println("conn.Read1 err=", err)
		return
	}
	fmt.Println("读取的buf长度为", buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据 pkgLen 读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read2 err=", err)
		return
	}
	//把pkgLen 反序列化成 -> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) fail", err)
		return
	}
	fmt.Printf("服务器，发送消息的长度=%d 内容=%s", len(data), string(data))

	n, err = conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	return
}
