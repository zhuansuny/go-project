package service
import (
	"go_code/chapter11_customerManage/model"
)

type CustomerService struct{
	customers []model.Customer
	customerNum int

}

func NewCustomerService()*CustomerService{
	cs := &CustomerService{}
	cs.customerNum =1
	customer :=model.NewCustomer(1,"张三","男",18,"112","@162.com")
	cs.customers =append(cs.customers,customer)
	return cs

} 

func(this *CustomerService) List() []model.Customer{
	return this.customers

}

func(this *CustomerService) Add(customer model.Customer) bool{
	this.customerNum ++
	customer.Id = this.customerNum
	this.customers =append(this.customers,customer)
	return true

}

func(this *CustomerService) FindById(id int)int{
	num :=-1
	for i, v := range this.customers {
		if v.Id == id{
		 num =i
		 break
		}
	}
	return num
}

func(this *CustomerService) Delete(id int) bool{
	index := this.FindById(id)
	if index == -1{
		return false
	}else{
		this.customers = append(this.customers[:index],this.customers[index+1:]...)
		return true
	}
	
}

func(this *CustomerService) Modify(id int,customer model.Customer) bool{
	index := this.FindById(id)
	if index == -1{
		return false
	}else{
		customers := append(this.customers[:index],customer)
		this.customers = append(customers[:index+1],this.customers[index+1:]...)
		return true
	}
	
}




