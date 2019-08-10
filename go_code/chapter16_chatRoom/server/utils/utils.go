package utils

//工具包，一些常用的函数、结构体，提供常用的方法
import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chapter16_chatRoom/common/message"
	"io"
	"net"
)

//将工具包的方法关联到结构体zhong
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //这时传输时，使用的缓存
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	//buf := make([]byte, 8096)
	fmt.Println("读取服务端发送的数据。。。")
	//conn.Read 只有在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭conn，就不会堵塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		if err == io.EOF {
			//fmt.Println("客户端退出，服务器协程也退出", err)
			return
		}
		fmt.Println("conn.Read1 err=", err)
		return
	}
	fmt.Println("读取的buf长度为", this.Buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//根据 pkgLen 读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read2 err=", err)
		return
	}
	//把pkgLen 反序列化成 -> message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) fail", err)
		return
	}
	fmt.Printf("服务端发送内容=%s", string(data))

	n, err = this.Conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	return
}
