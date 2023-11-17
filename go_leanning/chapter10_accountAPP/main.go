package main  //家庭收支软件普通方法实现
import (
	"fmt"
)



//主函数
func main(){

	
	var money float64
	var moneySum float64
	var str string
	var num int
	details := "收支\t账户金额\t收支金额\t说   明 \n"
	flag:
	for{

		fmt.Println("-------------------------家庭收支记账软件-----------------------")
		fmt.Println("                      1.收支明细")
		fmt.Println("                      2.登记收入")
		fmt.Println("                      3.登记支出")
		fmt.Println("                      4.退出软件")
		fmt.Scan(&num)
		switch num {
		case 1: 
			fmt.Println("-------------------------当前收支明细记录-----------------------")
			if details =="收支\t账户金额\t收支金额\t说   明 \n"{
				fmt.Println("当前没有明细，请去添加一笔吧")
				break
			}
			fmt.Println(details)
		
		case 2: 
			fmt.Println("-------------------------请登记收入金额-----------------------")
			fmt.Scan(&money)
			moneySum +=money
			fmt.Println("-------------------------请添加备注-----------------------")
			fmt.Scan(&str)
			details += fmt.Sprintf(" 收入\t     %v\t     %v \t    %v\n",moneySum,money,str)
			
		case 3: 
			fmt.Println("-------------------------请登记收入金额-----------------------")
			fmt.Scan(&money)
			if money>moneySum{
				fmt.Println("余额不足")
				break
			}
			moneySum -=money
			fmt.Println("-------------------------请添加备注-----------------------")
			fmt.Scan(&str)
			details += fmt.Sprintf(" 支出\t    %v\t     %v \t     %v\n",moneySum,money,str)
		case 4: 
			fmt.Println("你确定要退出吗？ y/n")
			choice :=""
			for{
				fmt.Scan(&choice)
				if choice=="y" || choice=="n"{
					break
				}
				fmt.Println("输入y退出软件，输入n返回")
			}
			if choice == "y"{
				break flag
			}
			
		default:
			fmt.Println("请重试")
			
			
		}

	}
	fmt.Println("家庭收支记账软件退出")

}
