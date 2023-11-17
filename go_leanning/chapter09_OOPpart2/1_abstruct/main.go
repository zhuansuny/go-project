package main  //编程思想，就是将现实中的事物抽象话，使用代码表达出来
import (
	"fmt"
)

type Account struct{    //银行账户
	accountNum string
	password string
	balance float64
}

func NewAccount(accountNum string,password string,balance float64)*Account{
	if len(accountNum)<6 || len(accountNum)>10{
		return nil
	}
	if balance <20{
		return nil
	}

	if len(password) !=6{
		return nil
	}
	return &Account{
		accountNum :accountNum ,
		balance :balance,
	 	password :password,
	}
}

func (account *Account) WithDraw(money float64,pwd string) {  //取款
	if pwd !=account.password{
		fmt.Println("密码错误")
		return
	}

	if(money>account.balance){
		fmt.Println("余额不足")
		return
	}else{
		account.balance -=money
		fmt.Println("取款成功")
	}
}


func (account *Account) Deposite(money float64,pwd string) { //存款
	if pwd !=account.password{
		fmt.Println("密码错误")
		return
	}
		account.balance +=money
		fmt.Println("存款成功")
}

func (account *Account) Query(pwd string) {   //查询
	if pwd !=account.password{
		fmt.Println("密码错误")
		return
	}
		fmt.Printf("您的账户%v 的余额为 %v\n",account.accountNum,account.balance)
}


//主函数
func main(){
	// var account = Account{        //创建一个账户
	// 	accountNum :"gs111111",
	// 	password : "666666",
	// 	balance :100,
	// }
	account := NewAccount("js1111","666666",100)  //创建一个账户

	if account == nil{
		fmt.Println("创建失败")
	}else{
		fmt.Println("创建成功",account)
	}
	var psw string 
	var money float64
	var num int
	flag:
	for{
		fmt.Println("输入数字确认使用的功能  1：取款 2：存款 3：查询余额 4：退出")
		fmt.Scan(&num)
		switch num {
		case 1: 
			fmt.Println("请输入取款金额")
			fmt.Scan(&money)
			fmt.Println("请输入密码")
			fmt.Scan(&psw)
			account.WithDraw(money,psw)
		case 2: 
			fmt.Println("请输入存款金额")
			fmt.Scan(&money)
			fmt.Println("请输入密码")
			fmt.Scan(&psw)
			account.Deposite(money,psw)
		case 3: 
			fmt.Println("请输入密码")
			fmt.Scan(&psw)
			account.Query(psw)
		case 4: 
			break flag
		default:
			fmt.Println("请重试")
			
			
		}

	}
	fmt.Println("系统退出")
	// account.Deposite(100,"666666")
	// account.Query("666666")
	// account.WithDraw(200,"666666")
	// account.Query("666666")

}
