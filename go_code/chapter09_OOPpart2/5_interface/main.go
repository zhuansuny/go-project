package main  //接口 高内聚低耦合

import (
	"fmt"
)

type USB interface { //定义一个接口
	//声明了两个没有实现的方法
	Start()                //实现这个接口需要实现这个接口的所有方法
	Stop()                 //接口里面的方法不可以有方法体
}

type wifi interface{
	Connect()
}


type object interface{} //没有任何方法的空接口，所有类型都实现了空接口

//定义一个手机结构体实现USB接口
type Phone struct{           //不需要指定实现USB接口，只要实现接口方法就实现了接口
 
}                            //可以实现多个接口

func (p Phone) Start(){
	fmt.Println("手机开始工作")
}

func (p Phone) Stop(){
	fmt.Println("手机停止工作")
}
func (p Phone) Connect(){            //实现了wifi接口 ，可以实现多个接口
	fmt.Println("手机连接了Wifi")
}

//定义一个相机结构体实现USB接口
type Camera struct{  

}

func (c Camera) Start(){
	fmt.Println("相机开始工作")
}

func (c Camera) Stop(){
	fmt.Println("相机停止工作")
}

//定义一个电脑结构体
type Computer struct{

}

func(c Computer) Working(usb USB){
	usb.Start()
	usb.Stop()
}


func main(){
	computer :=Computer{}
	phone :=Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)

	var usb USB = Phone{}  //把一个手机当成USB
	fmt.Println(usb)
	usb.Start()   //可以调用usb的方法，但不能调用手机的其他方法

}