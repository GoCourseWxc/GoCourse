package main

import (
	"fmt"
)

type FamilyAccount struct {
	// 声明一个字段，保存接收用户输入的选项
	key string
	// 声明一个字段，控制是否退出 for
	loop bool
	// 定义账户的余额 []
	balance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	// 定义字段，记录是否有收支的行为
	flag bool
	// 收支的详情使用字符串来记录
	// 当有收支时，只需要对 details 进行拼接处理即可
	details string
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户余额\t收支余额\t说  明",
	}
}

// 将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("========当前收支明细==========")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支密西")
	}
}

// 将登记收入写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money // 修改账户余额
	fmt.Println("本次收入说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 将登记支出写成一个方法，和 *FamilyAccount 绑定
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("账户余额不足")
	}
	this.balance -= this.money // 修改账户余额
	fmt.Println("本次支出说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		} else {
			fmt.Println("你输入的有误，请重新输入y/n")
		}
	}

	if choice == "y" {
		this.loop = false
	}
}

func (this *FamilyAccount) MainMenu() {
	// 显示这个主菜单
	for {
		fmt.Println("===========================家庭收支记账软件==============================")
		fmt.Println("===========================1.收支明细==============================")
		fmt.Println("===========================2.登记收入==============================")
		fmt.Println("===========================3.登记支出==============================")
		fmt.Println("===========================4.退出条件==============================")
		fmt.Println("请选择（1-4）")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确选项")
		}
		if !this.loop {
			break
		}
	}
}

func main() {
	NewFamilyAccount().MainMenu()
}
