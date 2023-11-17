package main//四种写文件的方式
import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func main(){

//---------------------方法一：创建一个文件并写入hello world----------------
	filePath :="d:/abc1.txt"  //文件路径
	file,err := os.OpenFile(filePath,os.O_WRONLY|os.O_CREATE,0666)  //打开文件

	if err !=nil {   //判断是否打开文件正常

		fmt.Println("打开文件错误：",err)
		return
	}
	defer file.Close()  //defer 一个文件关闭，在main函数执行完毕后释放资源
	str := "hello,world \r\n"  //定义要写入的字符串 \r\n可以在记事本中显示换行

	writer :=bufio.NewWriter(file)    //使用带缓存的*Writer
	for i := 0; i < 5; i++ {
		writer.WriteString(str)  //连续写入五次
		
	}
		//这时writer是在缓存区，没有真正的写入文件
		//需要Flush方法写入
	writer.Flush()

//---------------------方法二：打开一个存在的文件并清除写入hello golang----------------

	filePath1 :="d:/abc.txt"  //文件路径
	file1,err := os.OpenFile(filePath1,os.O_WRONLY|os.O_TRUNC,0666)  //打开一个存在的文件并清除写入

	if err !=nil {   //判断是否打开文件正常

		fmt.Println("打开文件错误：",err)
	}
	defer file1.Close()  //defer 一个文件关闭，在main函数执行完毕后释放资源
	str = "hello,golang \r\n"  //定义要写入的字符串 \r\n可以在记事本中显示换行

	writer1 :=bufio.NewWriter(file1)    //使用带缓存的*Writer
	for i := 0; i < 5; i++ {
		writer1.WriteString(str)  //连续写入五次
		
	}
		//这时writer是在缓存区，没有真正的写入文件
		//需要Flush方法写入
	writer1.Flush()
//---------------------方法三：打开一个存在的文件并追加写入hello file----------------

	//filePath ="d:/abc.txt"  //文件路径
	file2,err := os.OpenFile(filePath,os.O_WRONLY|os.O_APPEND,0666)  //打开一个存在的文件并追加

	if err !=nil {   //判断是否打开文件正常

		fmt.Println("打开文件错误：",err)
	}
	defer file2.Close()  //defer 一个文件关闭，在main函数执行完毕后释放资源
	str = "hello file\r\n"  //定义要写入的字符串 \r\n可以在记事本中显示换行

	writer2 :=bufio.NewWriter(file2)    //使用带缓存的*Writer
	for i := 0; i < 5; i++ {
		writer2.WriteString(str)  //连续写入五次
		
	}
		//这时writer是在缓存区，没有真正的写入文件
		//需要Flush方法写入
	writer2.Flush()

	//---------------------方法四：打开一个存在的文件读取数据并追加写入hello file----------------

	//filePath ="d:/abc.txt"  //文件路径
	file3,err := os.OpenFile(filePath1,os.O_RDWR|os.O_APPEND,0666)  //打开一个存在的文件读取数据并追加写入

	if err !=nil {   //判断是否打开文件正常

		fmt.Println("打开文件错误：",err)
	}
	
	defer file3.Close()  //defer 一个文件关闭，在main函数执行完毕后释放资源
	str = "hello 你好\r\n"  //定义要写入的字符串 \r\n可以在记事本中显示换行
	reader := bufio.NewReader(file3)
	for {
		str5 ,err :=reader.ReadString('\n')  //读到一个换行就结束
		if err ==io.EOF {
			break
		}

		fmt.Print(str5)
	}
	writer3 :=bufio.NewWriter(file3)    //使用带缓存的*Writer
	for i := 0; i < 5; i++ {
		writer3.WriteString(str)  //连续写入五次
		
	}
		//这时writer是在缓存区，没有真正的写入文件
		//需要Flush方法写入
	writer3.Flush()

}