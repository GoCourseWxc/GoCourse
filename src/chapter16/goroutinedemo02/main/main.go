package main

import (
	"fmt"
	"time"
)

// 计算 1 - 200 的各个数的阶乘，并且把各个数的阶乘放入到map中
//
var (
	myMap = make(map[int]int, 10)
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}

func main() {
	for i := 0; i <= 10; i++ {
		go test(i)
	}
	time.Sleep(10 * time.Second)
	for k, v := range myMap {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
}
