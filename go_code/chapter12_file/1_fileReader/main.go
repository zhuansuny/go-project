package main  //文件是用来保存数据的


//流：分为输入流和输出流 是数据在数据源（文件）和程序（内存）之间经历的路径
//在golang中由File结构体封装了文件的所有操作来实现（os包）
import (
	"fmt"
	"os"
	"bufio"
	"io"
	"io/ioutil"
)


func main(){
	//---------------------------带缓冲的读取方式（文件较大时）-----------------------------------------
	file,err :=os.Open("d:/lol.txt")        //file叫文件对象/指针/文件句柄
	if err !=nil{
		fmt.Println("open file err=",err)
	}else{
		fmt.Printf("file =%v\n",file)
	}

	defer file.Close()
		
	reader := bufio.NewReader(file)

	for {
		str ,err :=reader.ReadString('\n')  //读到一个换行就结束
		if err ==io.EOF {
			break
		}

		fmt.Print(str)
	}
		
	fmt.Println("文件读取完毕")


	//---------------------------一次性的读取方式（文件较小时）-----------------------------------------
	file1 := "d:/lol2.txt"
	content ,err := ioutil.ReadFile(file1)
	if err !=nil{
		fmt.Println("read file err=",err)
	}
	//把读取的内容显示到终端
	fmt.Printf("%v",string(content))  //[]byte切片

	//没有显式的open文件，因此也不需要显式的close文件
	//因为文件的open和close被封装到ReadFile函数内部
	

}