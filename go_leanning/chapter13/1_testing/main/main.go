package main //单元测试

import (
	"fmt"
)

func addUpper(n int) int {
	res := 0
	for i := 1; i < n; i++ {
		res += i
	}
	return res
}

func getSub(n1 int, n2 int) int {
	return n1 - n2
}

func main() {
	res := addUpper(10)
	res2 := getSub(100, 77)
	fmt.Println(res, res2)

}
