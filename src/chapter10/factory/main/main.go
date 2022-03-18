package main

import (
	"GoCourse/src/chapter10/factory/model"
	"fmt"
)

func main() {
	// 通过工厂模式构造方法来访问对象
	student := model.NewStudent("tom", 88.0)

	fmt.Println(*student)
	fmt.Println("name=", student.Name, "score=", student.GetScore())
}
