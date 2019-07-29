package utils

import (
	"fmt"
)

type FamilyAccount struct{
	money float64
	moneySum float64
	str string
	num int
	details string 
}

func NewFamilyAccount() *FamilyAccount{
	return &FamilyAccount{
		details : "收支\t账户金额\t收支金额\t说   明 \n",
		moneySum : 1000,
	}
} 

func (this *FamilyAccount) showDetail(){
	fmt.Println("-------------------------当前收支明细记录-----------------------")
			if this.details =="收支\t账户金额\t收支金额\t说   明 \n"{
				fmt.Println("当前没有明细，请去添加一笔吧")
			}else{
				fmt.Println(this.details)
			}
}

func (this *FamilyAccount) income(){
	fmt.Println("-------------------------请登记收入金额-----------------------")
	fmt.Scan(&this.money)
	this.moneySum +=this.money
	fmt.Println("-------------------------请添加备注-----------------------")
	fmt.Scan(&this.str)
	this.details += fmt.Sprintf(" 收入\t     %v\t     %v \t    %v\n",this.moneySum,this.money,this.str)

}
func (this *FamilyAccount) expend(){
	fmt.Println("-------------------------请登记支出金额-----------------------")
	fmt.Scan(&this.money)
	if this.money>this.moneySum{
		fmt.Println("余额不足")
		return
	}
	this.moneySum -=this.money
	fmt.Println("-------------------------请添加备注-----------------------")
	fmt.Scan(&this.str)
	this.details += fmt.Sprintf(" 支出\t    %v\t     %v \t     %v\n",this.moneySum,this.money,this.str)
}

func (this *FamilyAccount) exit()string{
	fmt.Println("你确定要退出吗？ y/n")
	choice :=""
	for{
		fmt.Scan(&choice)
		if choice=="y" || choice=="n"{
			break
		}
		fmt.Println("输入y退出软件，输入n返回")
	}
	return choice
}

func (this *FamilyAccount) MainMenu(){
	flag:
	for{

		fmt.Println("-------------------------家庭收支记账软件-----------------------")
		fmt.Println("                      1.收支明细")
		fmt.Println("                      2.登记收入")
		fmt.Println("                      3.登记支出")
		fmt.Println("                      4.退出软件")
		fmt.Scan(&this.num)
		switch this.num {
		case 1: 
			this.showDetail()
		
		case 2: 
		    this.income()
			
		case 3: 
			this.expend()
		case 4: 
			choice:=this.exit()
			if choice == "y"{
				break flag
			}
			
		default:
			fmt.Println("请重试")
			
			
		}

	}
	fmt.Println("家庭收支记账软件退出")

}