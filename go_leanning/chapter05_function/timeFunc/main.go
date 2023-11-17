package main
import (
	"fmt"
	"time"
)

func main(){
   //获取当前时间
	now := time.Now()
	fmt.Printf("now的值%v now的类型%T",now,now)

   //2.通过now可以获取到年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期时间

	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), 
	now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), 
	now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dateStr=%v\n", dateStr)

	//格式化日期时间的第二种方式
	fmt.Printf(now.Format("2006-01-02 15:04:05"))  //必须要写这个时间
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	fmt.Printf(now.Format("2006"))
	fmt.Println()


	time.Sleep(time.Millisecond * 100)  //休眠100毫秒
	//Unix（秒）和UnixNano（纳秒）当前时间（从1970年1月1号8点）的使用时间戳
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())

}
	

