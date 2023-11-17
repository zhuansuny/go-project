package utils
import (
	"fmt"
)
func Cal()  {                   //同一个包下不可以取相同的函数名
	fmt.Println("调用Cal成功")
}                    //可以生成一个utils.a的库文件，可以不用源码就可以调用包内函数（一般会有说明文档）