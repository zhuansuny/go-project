package main
import (
	"fmt"

	"io/ioutil"
)

func main(){
	//将d:/abc.txt导入到c：/kkk.txt
	//1首先将d:/abc.txt内容读取到内存
	//2将读取的内容写入到c：/kkk.txt
	filePath :="d:/abc.txt"
	file2Path :="d:/kkk.txt"

	data ,err :=ioutil.ReadFile(filePath)
	if err !=nil {

		fmt.Println("读取文件错误",err)
		return
	}
	fmt.Println("读取文件成功")
	err = ioutil.WriteFile(file2Path,data,0666)
	if err !=nil {
		fmt.Println("写入文件错误",err)
		return
	} 
	fmt.Println("写入文件成功")

}
