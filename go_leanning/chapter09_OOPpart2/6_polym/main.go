package main //面向对象三大特性之多态
			 //golang中多态是由接口实现
import (
	"fmt"
)
//-------------------------USB接口--------------------------
type USB interface { 
	Start()                
	Stop()                
}

//-------------------------phone结构体--------------------------

type Phone struct{           //不需要指定实现USB接口，只要实现接口方法就实现了接口
 	Name string
}                            //可以实现多个接口
	
func (p Phone) Start(){
	fmt.Println(p.Name,"手机USB开始工作")
}
	
func (p Phone) Stop(){
	fmt.Println(p.Name,"手机USB停止工作")
}
func (p Phone) Call(){
	fmt.Println(p.Name,"手机开始呼叫")
}

//-------------------------camera结构体--------------------------
type Camera struct{  
	Name string
}

func (c Camera) Start(){
	fmt.Println(c.Name,"相机USB开始工作")
}

func (c Camera) Stop(){
	fmt.Println(c.Name,"相机USB停止工作")
}

//-------------------------computer结构体--------------------------
type Computer struct{

}

func(c Computer) Working(usb USB){  //把USB接口当作参数
	p ,ok :=usb.(Phone)
	if ok {         //进行类型断言判断是否是phone，进行特有的方法
		p.Start()
		p.Call()
		p.Stop()
		return
	}

	ca ,ok := usb.(Camera)
	if ok {
		ca.Start()
		ca.Stop()
	}else{
		fmt.Println("convert fail")
	}
	
}


func main(){
//接口体现多态的两种类型
//-------------------------1.多态参数--------------------------

	computer :=Computer{}
	phone :=Phone{"vivo"}
	camera := Camera{}

	computer.Working(phone)   //把USB当做参数，所有实现了USB接口的类型都可传入
	computer.Working(camera)

	var usb USB = Phone{"苹果"}  //把一个手机当成USB
	fmt.Println(usb)
	usb.Start()   //可以调用usb的方法，但不能调用手机的其他方法
	
//----------------------断言------------------------------------

	phone = usb.(Phone)   //类型断言  ,可以转成之前转换的Phone类型  必须清楚知道USB之前是什么类型转换的
	//camera =usb.(Camera)   //会报错，usb是phone转换的，不能再转换为camera ，并会退出程序
	camera , ok := usb.(Camera)  //可以进行类型断言检测 ,让程序继续执行
	if ok {
		fmt.Println("convert success")
		fmt.Printf("camera的类型是%T 值是 %v",camera,camera)
	}else{
		fmt.Println("convert fail")
	}


	phone.Call()      //苹果手机的呼叫


//--------------------2.多态数组-------------------------------
	//定义一个USB数组，可以存放phone以及camera
	var usbarr [3]USB
	usbarr[0] =Phone{"华为"}
	usbarr[1] =Phone{"小米"}
	usbarr[2] =Camera{"尼康"}
	fmt.Println(usbarr)


}