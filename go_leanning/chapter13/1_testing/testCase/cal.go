package cal

//一个被测试函数

//testing框架会将以_test.go结尾的文件引入import
//然后将Test开头的函数引入main
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

//求两个数的查
func getSub(n1 int, n2 int) int {
	return n1 - n2

}
