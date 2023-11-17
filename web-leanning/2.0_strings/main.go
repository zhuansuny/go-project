package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "fa"))  //false
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ",")) //用,把切片中的值拼接起来  foo,bar,baz

	fmt.Println("ba" + strings.Repeat("na", 10)) //重复s字符串count次，最后返回重复的字符串 banananananananananana

	//func Replace(s, old, new string, n int) string
	//在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println(strings.Replace("olink olink olink", "k", "ky", 2))   //olinky olinky olink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo

	// func Split(s, sep string) []string
	// 把s字符串按照sep分割，返回slice
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! ")) //在s字符串的头部和尾部去除cutset指定的字符串

	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   ")) //去除s字符串的空格符，并且按照空格分割返回slice

	//-------------------------strconv---------------------
	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	//Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	//Parse 系列函数把字符串转换为其他类型
	a1, err := strconv.ParseBool("false")
	checkError(err)
	b1, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c1, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d1, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e1, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a1, b1, c1, d1, e1)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
