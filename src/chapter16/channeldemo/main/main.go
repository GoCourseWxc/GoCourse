package main

import "fmt"

func main() {
	// 演示管道的使用
	// 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	// 查看intChan
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p \n", intChan, &intChan)

	// 管道内写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 管道写入数据，不能超过容量

	// 查看管道的长度和容量
	fmt.Printf("channel len= %v cap= %v \n", len(intChan), cap(intChan))

	// 从管道中读取数据
	num2 := <-intChan
	fmt.Printf("num2= %v\n", num2)
	// 在没有使用协程的情况下，如果管道数据已经全部取出，再取出就会报告 deadlock

	num3 := <-intChan
	fmt.Printf("num3= %v\n", num3)
	num4 := <-intChan
	fmt.Printf("num4= %v\n", num4)
}
