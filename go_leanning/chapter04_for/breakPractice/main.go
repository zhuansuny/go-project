package main
import(
	"fmt"
	
)

func main(){
	//总数num，每经过一次路口当大于50000时减少百分之五，小于50000减少1000，计算一共可以经过多少路口
	var num float64 = 100000000000
	var count int
	for i := 0; i < 100000; i++ {
		if(num<1000){
			break
		}
		if(num>50000){
			num = 0.95*num
		}else{
			num = num -1000
		}
		count++

	}	
	fmt.Println("一共可以经过路口次数为：",count)
}