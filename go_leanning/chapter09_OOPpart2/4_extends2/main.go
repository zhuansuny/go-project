package main   //面向对象三大特性之继承
import (
	"fmt"
)

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand	
}

type TV2 struct {
	*Goods     //继承可以使用指针，改成地址传输快
	*Brand	
}

func main(){
	tv1 :=TV{Goods{"黑白电视",2000.0},Brand{"格力","珠海"}}
	var tv2 = TV{
		Goods{
			Name : "液晶电视",
			Price :5000.0,
		},
		Brand{
			Name :"海尔",
			Address :"北京",
		},
	}

	fmt.Println("tv1",tv1)
	fmt.Println("tv2",tv2)



	tv3 :=TV2{&Goods{"黑白电视",2000.0},&Brand{"格力","珠海"}}  //嵌套结构体指针
	var tv4 = TV2{
		&Goods{
			Name : "液晶电视",
			Price :5000.0,
		},
		&Brand{
			Name :"海尔",
			Address :"北京",
		},
	}

	fmt.Println("tv3",*tv3.Goods,*tv3.Brand)
	fmt.Println("tv4",*tv4.Goods,*tv4.Brand)
}
