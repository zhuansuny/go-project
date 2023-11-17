package main  //map练习

import (
	"fmt"
)

func modifyUser(users map[string]map[string]string,name string ){
	if users[name] !=nil{
		users[name]["pws"] = "888888"
	}else{
		users[name] = make(map[string]string,2)
		users[name]["pws"] = "11111"
		users[name]["nickname"] ="昵称"+name
	}


}

func main(){
	users := make(map[string]map[string]string,10) 

	modifyUser(users,"tom")
	modifyUser(users,"mary")
	modifyUser(users,"jack")
	fmt.Println(users)
	modifyUser(users,"tom")
	fmt.Println(users)
}