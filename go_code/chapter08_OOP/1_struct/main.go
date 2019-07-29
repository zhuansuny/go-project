package main  //结构体是值类型
import (
	"fmt"
)
//golang 不是纯粹的面向对象语言，只能说支持面向对象编程特性 （面向接口编程）
//结构体相当于JAVA中的类
type cat struct{     //cat的首字母代表作用域
	name string    //name 叫做字段/属性/field
	age int			//字段可以是基本类型和引用类型
	color string
	// slice []int
	// map1 map[int]string
}

func main(){
	var cat3 cat                   //创建cat对象/实例
	cat1 :=cat{"小白",3,"白色"}  
	cat2 :=cat{"小花",10,"花色"}
	cat3 =cat{"小花",10,"花色"}
	cat3.name ="小红"
	// cat3.slice =make([]int,3)             //结构体中的切片和Map必须先make分配空间才可以使用
	// cat3.map1 =make(map[int]string,3)
	// cat3.slice[0] = 3
	// cat3.map1[0] = "3"
	fmt.Println(cat1,cat2,cat3)

	cat1 = cat3     //是值传递，改变cat1不会影响cat3的值
	cat1.name ="小岚"
	fmt.Println(cat1,cat3)

	//创建结构体的四种方式

	//方式一
	// var cat1 cat

	//方式二
	//cat1 :=cat{}   cat1 :=cat{"小白",3,"白色"}

	//方式三
	//var cat1 *cat = new(cat)    赋值 (*cat).name = "小红" 可以(底层有优化)省略为 cat.name =  "小红" 

	//方式四
	//var cat1 *cat = &cat{}   赋值同方式三




	


}