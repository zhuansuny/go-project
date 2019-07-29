package main  //map是引用类型

import (
	"fmt"
)

func main(){
	var m map[int]string    //map声明，声明不会分配内存，需要make来分配
	m = make(map[int]string,2)  //必须make才可以使用

	var m1 = make(map[int]string,10) //可以直接make
	fmt.Println(m1)

	m2 := map[int]string{     //自动赋值
		1 : "卢俊义",
	}
	m[1] = "宋江"
	m[2] = "吴用"
	m[1] = "武松"          //KEY相同会覆盖之前的值 ，value可以重复
	m[3] = "卢俊义"        //map可以自己动态增加
	fmt.Println(m)
	fmt.Println(m2)  
	fmt.Printf("m有%d对\n",len(m))

	delete(m,3)           //按key删除，没有该数据就不执行，不会报错
	fmt.Println(m)        //map完全无序不按插入排序也不按key大小排序

	val,ok := m[1]      //查找
	if ok{
		fmt.Println("存在值为：",val)
	}else{
		fmt.Println("不存在值")
	}
	m2 = make(map[int]string,10)  //golang没有全部清空函数，可以分配新的空间，旧的gc回收
	fmt.Println(m2) 

	//map的遍历（只能是for-range）
	for k, v := range m {
		fmt.Printf("k=%v v=%v\n",k,v)
	}




	//把map再作为value案例

	studentMap :=make(map[string]map[string]string)
	studentMap["No1"] = make(map[string]string,3)
	studentMap["No1"]["name"] ="tom"
	studentMap["No1"]["sex"] ="男"
	studentMap["No1"]["address"] ="北京"

	studentMap["No2"] = make(map[string]string,3)
	studentMap["No2"]["name"] ="mary"
	studentMap["No2"]["sex"] ="女"
	studentMap["No2"]["address"] ="上海"
	fmt.Println(studentMap["No1"])

	//map的遍历（只能是for-range）
	for k, v := range studentMap {
		fmt.Printf("k=%v v=%v\n",k,v)
		for k1, v1 := range v {
			fmt.Printf("k1=%v v1=%v\n",k1,v1)
		}
		
	}
    
	
}