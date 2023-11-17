package main  //string字符串常用的系统函数
import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	str := "你好a"    //一个中文字符占3个字节（utf—8）数字/字母一个字节
	fmt.Println(len(str))  //len(str)统计字符串字节长度  结果7

	str2 := "hello北京"
	r := []rune(str2)   //切片函数，将字符串按照字符输出 ，处理含有中文的字符串遍历问题
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符= %c\n",r[i])
	}


	n ,err := strconv.Atoi("123")  //字符串转整数  err接受错误
	if err !=nil{
		fmt.Println("转换错误",err)
	}else{
		fmt.Println("转换结果是：",n)
	}

	str =strconv.Itoa(12345)  //整数转字符串
	fmt.Println(str)
	
	var bytes = []byte("hello go") //字符串转[]byte
	fmt.Println(bytes)

	str = string([]byte{97,98,99})  //[]byte转字符串
	fmt.Println(str)

	str = strconv.FormatInt(123,2) //将十进制123转换为2进制返回为字符串
	fmt.Println(str)
	str = strconv.FormatInt(123,16) //将十进制123转换为16进制返回为字符串
	fmt.Println(str)

	//查找子串是否在指定的字符串中: strings.Contains("seafood", "foo") //true
	b := strings.Contains("seafood", "mary")
	fmt.Printf("b=%v\n", b) 

	//统计一个字符串有几个指定的子串 ： strings.Count("ceheese", "e") //4
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v\n", num)

	//10)不区分大小写的字符串比较(==是区分字母大小写的): fmt.Println(strings.EqualFold("abc", "Abc")) // true

	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b) //true

	fmt.Println("结果","abc" == "Abc") // false //区分字母大小写

	//11)返回子串在字符串第一次出现的index值，如果没有返回-1 : 
	//strings.Index("NLT_abc", "abc") // 4

	index := strings.Index("NLT_abcabcabc", "abc") // 4
	fmt.Printf("index=%v\n",index)

	//12)返回子串在字符串最后一次出现的index，
	//如没有返回-1 : strings.LastIndex("go golang", "go")

	index = strings.LastIndex("go golang", "go") //3
	fmt.Printf("index=%v\n",index)

	//将指定的子串替换成 另外一个子串: strings.Replace("go go hello", "go", "go语言", n) 
	//n可以指定你希望替换几个，如果n=-1表示全部替换

	str2 = "go go hello"
	str = strings.Replace(str2, "go", "北京", -1)
	fmt.Printf("str=%v str2=%v\n", str, str2)

	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组： 
	//strings.Split("hello,wrold,ok", ",")
	strArr := strings.Split("hello,wrold,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	} 
	fmt.Printf("strArr=%v\n", strArr)

	//15)将字符串的字母进行大小写的转换: 
	//strings.ToLower("Go") // go strings.ToUpper("Go") // GO

	str = "goLang Hello"
	str = strings.ToLower(str) 
	str = strings.ToUpper(str) 
	fmt.Printf("str=%v\n", str) //golang hello

	//将字符串左右两边的空格去掉： strings.TrimSpace(" tn a lone gopher ntrn   ")
	str = strings.TrimSpace(" tn a lone gopher ntrn   ")
	fmt.Printf("str=%q\n", str)

	//17)将字符串左右两边指定的字符去掉 ： 
	//strings.Trim("! hello! ", " !")  // ["hello"] //将左右两边 ! 和 " "去掉
	str = strings.Trim("! he!llo! ", " !")
	fmt.Printf("str=%q\n", str)

	//20)判断字符串是否以指定的字符串开头: 
	//strings.HasPrefix("ftp://192.168.10.1", "ftp") // true

	b = strings.HasPrefix("ftp://192.168.10.1", "hsp") //true
	fmt.Printf("b=%v\n", b)
	

}