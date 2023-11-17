package main
import (
	"fmt"
)
func test(n int){   //递归是函数内调用本身
	if(n>2){        //递归必须是向退出递归条件逼近，否则是死循环
		n--
		test(n)
		//return   //遇到return函数会立即返回，遵守谁调用将结果返回给谁，当函数执行完毕或者返回时，函数会销毁
	}
	fmt.Println(n)
	
}

func fb(n int)int{    //斐波那契数计算 1 ,1 ,3 ,5 ,8 ,13 当前数是前面两个数之和
	if n==1||n==2{
		return 1
	}
	num := fb(n-1)+fb(n-2)
	return num
}

func f(n int) int{      //实现 f(n)=2*f(n-1)+1 函数
	if n==1{
		return 3
	}else{
		num := 2*f(n-1)+1
		return num
	}
}


func peach(day int) int{  //猴子吃桃问题
	if(day==1){           //每天吃前一天总数的一半加1，问：吃到第十天时只剩一个，总共有多少个桃
		return 1
	}
	num := 2 + 2*peach(day-1) 
	return num

}
func main(){
	fmt.Println(fb(7))
	fmt.Println(f(5))
	fmt.Println(peach(10))  //吃桃结果为1534
}