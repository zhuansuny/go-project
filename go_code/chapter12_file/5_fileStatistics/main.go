package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)
type CharCount struct{
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}

func main(){
	//打开一个文件，创建一个Reader
	//每读一行，就去统计该行有多少个 英文
	//然后将结果保存到一个结构体中
	fileName := "d:/abc.txt"   
	file,err := os.Open(fileName)
	if err !=nil{
		fmt.Println("open file err=",err)
	}
	defer file.Close()
	var count CharCount

	reader := bufio.NewReader(file)
	coun:=0

	for {
		str,err :=reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//str1 := []rune(str)
		
		for _, v := range str {
			fmt.Println(string(v))
			coun ++
			
			switch {                     //switch 不添加条件可以当作if else使用
			case v>='a'&& v<='z':
				fallthrough  //穿透
			case v>='A'&& v<='Z':
				count.ChCount ++

			case v==' '|| v=='\t':
				count.SpaceCount ++

			case v>='0'&& v<='9':
				count.NumCount ++
			default :
				count.OtherCount ++
				
			}
		}
		
	}
	fmt.Println("统计",coun)
	fmt.Println(count)

}
