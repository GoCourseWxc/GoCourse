package main

import (
	model "GoCourse/src/chapter13/model"
	"GoCourse/src/chapter13/service"
	"fmt"
)

type customerView struct {
	// 定义必要字段
	key  string // 接收用户输入...
	loop bool   // 表示是否循环的显示主菜单

	// 增加一个字段 customerService
	customerService *service.CustomerService
}

func (this *customerView) list() {
	// 获取当前所有客户信息

	customers := this.customerService.List()
	fmt.Println("===========================4.客户列表==============================")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
}

// 添加用户
func (this *customerView) add() {
	fmt.Println("===========================1.添加客户==============================")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)

	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("邮件：")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("===========================添加完成==============================")
	} else {
		fmt.Println("===========================添加失败==============================")
	}
}

// 得到用户的输入id，删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("===========================3.删除客户==============================")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	if id == -1 {
		return // 放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N):")
	choice := ""
	fmt.Println(choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("===========================刪除完成==============================")
		} else {
			fmt.Println("===========================刪除失败输入的id不存在==============================")
		}
	}
}

func (this *customerView) MainMenu() {

	// 显示这个主菜单
	for {

		fmt.Println("===========================客户信息管理软件==============================")
		fmt.Println("===========================1.添加客户==============================")
		fmt.Println("===========================2.修改客户==============================")
		fmt.Println("===========================3.删除客户==============================")
		fmt.Println("===========================4.客户列表==============================")
		fmt.Println("===========================5.退出==============================")
		fmt.Println("请选择（1-5）")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.add()
		case "2":
			//
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.loop = false
		default:
			fmt.Println("请输入正确选项")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统")
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.MainMenu()
}
