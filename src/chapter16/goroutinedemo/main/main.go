package main

import (
	"fmt"
	"strconv"
	"time"
)

// 在主线程中，开启一个goroutine,该协程每隔1秒输出 "helo,world"
// 在主线程中每隔一秒输出 “hello，golang",输出10次后，退出程序
// 要求主线程和goroutine同时执行

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	go test() // 开启一个协程

	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
