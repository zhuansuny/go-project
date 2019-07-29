package main    //

import (
	"fmt"
	"go_code/chapter11_customerManage/service"
	"go_code/chapter11_customerManage/model"
)

type customerView struct{
 
	key int                     //switch判断条件
	loop bool                   //退出for循环的标志
	customerService *service.CustomerService    //用来调用customerService的方法

}

func (this *customerView) list(){
	customs := this.customerService.List()  //调用customerService的List方法
	fmt.Println("----------------------------客户列表--------------------")

	fmt.Println("编号\t 姓名\t 性别\t 年龄 \t电话\t 邮箱")
	
	for i := 0; i < len(customs); i++ {
		fmt.Println(customs[i].GetInfo())   //对custom输出格式化
	}


	fmt.Println("----------------------------客户列表完成--------------------")
}


func  (this *customerView) add(){
	fmt.Println("-----------请添加客户--------------")
	fmt.Println("姓名：")
	name :=" "
	fmt.Scanln(&name)

	fmt.Println("性别：")
	gender :=" "
	fmt.Scanln(&gender)

	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)

	fmt.Println("电话：")
	phone :=" "
	fmt.Scanln(&phone)

	fmt.Println("电邮：")
	email :=" "
	fmt.Scanln(&email)
	custom := model.NewCustomer2(name,gender,age,phone,email)
	if this.customerService.Add(custom){
		fmt.Println("添加完成")
	}else{
		fmt.Println("添加失败")
	}

}

func (this *customerView)delete(){
	fmt.Println("-----------删除客户--------------")
	fmt.Println("请输入要删除客户编号（-1退出）")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return
	}
	fmt.Println("确认是否删除（y/n）")
	choice :=""
	fmt.Scanln(&choice)
	if choice=="y"||choice=="Y"{
		if this.customerService.Delete(id) {
			fmt.Println("删除完成")
		}else{
			fmt.Println("删除失败，输入的ID号不存在")
		}
	}
}

func (this *customerView)modify(){
	fmt.Println("-----------修改客户--------------")
	fmt.Println("请输入要修改的客户编号（-1退出）")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return
	}
	fmt.Println("确认是否修改（y/n）")
	choice :=""
	fmt.Scanln(&choice)

	if choice=="y"||choice=="Y"{
		fmt.Println("-----------输入修改的客户信息--------------")
		fmt.Println("姓名：")
		name :=" "
		fmt.Scanln(&name)

		fmt.Println("性别：")
		gender :=" "
		fmt.Scanln(&gender)

		fmt.Println("年龄：")
		age := 0
		fmt.Scanln(&age)

		fmt.Println("电话：")
		phone :=" "
		fmt.Scanln(&phone)

		fmt.Println("电邮：")
		email :=" "
		fmt.Scanln(&email)
		custom := model.NewCustomer(id,name,gender,age,phone,email)
		if this.customerService.Modify(id,custom) {
			fmt.Println("修改完成")
		}else{
			fmt.Println("修改失败，输入的ID号不存在")
		}
	}
}


func (this *customerView) exit(){
	choice :=""
	for{
		fmt.Println("确认是否退出（y/n）")
		fmt.Scanln(&choice)
		if choice=="y"||choice=="Y"||choice=="n"||choice=="N"{	
			break
		}
	}
	if choice=="y"||choice=="Y"{
		this.loop = !this.loop
	}
}
	
func (this *customerView) mainMenu(){
	for !this.loop {
		fmt.Println("-------------------客户信息管理软件----------------------")
		fmt.Println("                      1.添加客户")
		fmt.Println("                      2.修改客户")
		fmt.Println("                      3.删除客户")
		fmt.Println("                      4.客户列表")
		fmt.Println("                      5.退出软件")
		fmt.Println("请选择（1-5）：")

		fmt.Scanln(&this.key)
		
	
		switch this.key {
			case 1: 
				this.add()
			case 2: 
				this.modify()
				
			case 3: 
				this.delete()
			case 4: 
				this.list()
			case 5: 
				this.exit()
			default:
				fmt.Println("请重试")			

		}

	}
	fmt.Println("客户信息管理软件退出")

}


//主函数
func main(){
	var cv customerView
	cv.customerService =service.NewCustomerService()   //首先创建一个customerService实例
	cv.mainMenu()
}