package main//解析CMD中命令行的数据

import (
	"fmt"
	"os"
	"flag"
)

func main(){
	//--------1.原生的解析命令行参数------------
	fmt.Println("命令行的参数有",len(os.Args))  //直接调用os.Args可以得到
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n",i,v)
	}
	//--------1.flag包解析命令行参数------------    

	//可以不用按顺序输入参数  ，通过-u  -pwd   -h  -port 后面加上参数 可以不用按照顺序
	 var user string
	 var pwd string
	 var host string
	 var port int

	 flag.StringVar(&user,"u","","用户名默认为空")
	 flag.StringVar(&pwd,"pwd","","密码默认为空")
	 flag.StringVar(&host,"h","localhost","主机名默认为localhost")
	 flag.IntVar(&port,"port",3306,"端口名默认为3306")

	 flag.Parse()

	 fmt.Printf("user=%v pwd=%v  host=%v  port=%v",user,pwd,host,port)
	
}