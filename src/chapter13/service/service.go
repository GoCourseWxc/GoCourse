package service

import "GoCourse/src/chapter13/model"

type CustomerService struct {
	customers []model.Customer
	//
	customerNum int
}

func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "123", "812312@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

// 根据id查看客户在切片下标，如果没有该客户，返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	// 遍历this.customers 切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = 1
		}
	}
	return index
}

// 删除用户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
