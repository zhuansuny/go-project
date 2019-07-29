package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

//
func CopyFile(dstFile string ,srcFile string) (written int64 ,err error){
	srcfile,err :=os.Open(srcFile)
	if err !=nil{
		fmt.Println("读取文件错误",err)
	}
	reader :=bufio.NewReader(srcfile)
	defer srcfile.Close()

	dstfile,err := os.OpenFile(dstFile,os.O_WRONLY|os.O_CREATE,0666)  //打开文件

	if err !=nil {   //判断是否打开文件正常

		fmt.Println("打开文件错误：",err)
		return
	}
	defer dstfile.Close() 
	writer :=bufio.NewWriter(dstfile)
	return io.Copy(writer,reader)   // 调用io包里的Copy函数  返回copy的字节数  需要传入文件的Writer和Reade实例
}

func main(){
	//将d:/flower.jpg导入到d：/abc.jpg
	srcFile := "d:/flower.jpg"   //copy的图片文件目录
	dstFile := "d:/abc.jpg"       //copy到的目录，如果不存在将会创建一个
	_,err := CopyFile(dstFile,srcFile)  //调用自定义的函数将src以及dst输入
	if err != nil{
		fmt.Println("copy文件错误：",err)
		return
	}
	fmt.Println("copy文件完成")

}
