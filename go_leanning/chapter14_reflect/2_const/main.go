package main //const常量

import (
	"fmt"
)

func main() {
	var num int
	//常量定义时必须赋值
	const tax int = 0
	//常量不可以被修改
	//tax = 10 错误
	const num2 = 3.0 / 9
	fmt.Println(num, tax, num2)

	const ( //简易定义方法
		a = iota //a=0
		b        //下面每一行递增1
		c
		d
		e, f = iota, iota //ef都为4
	)
	fmt.Println(a, b, c, d, e, f)
}
