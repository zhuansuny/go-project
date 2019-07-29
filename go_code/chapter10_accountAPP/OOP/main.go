package main  //家庭收支软件面向对象方法实现
import (
	"fmt"
	"go_code/chapter10_accountAPP/OOP/utils"
)



//主函数
func main(){
	fmt.Println("这个是面向对象的方式完成")
	utils.NewFamilyAccount().MainMenu()   //通过utils包的NewFamilyAccount函数创建一个FamilyAccount实例 
	                                      //然后调用MainMenu方法进行实现功能
}
