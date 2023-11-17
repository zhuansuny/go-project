package main  //map结构体

import (
	"fmt"
)

type stu struct{  //结构体定义
	name string
	age int
	addrss string
}

func main(){

	
	student := make(map[int]stu,2)  //必须make才可以使用
	stu1 := stu{"tom",18,"北京"}
	stu2 := stu{"mary",17,"上海"}
	student[0]=stu1
	student[1]=stu2
	fmt.Println(student)

	for k, v := range student {
		fmt.Printf("学生的学号是%d\n",k)
		fmt.Printf("学生的名字是%v\n",v.name)   //map结构体遍历
		fmt.Printf("学生的年龄是%v\n",v.age)
		fmt.Printf("学生的地址是%v\n",v.addrss)
		fmt.Println()

	}
	
}