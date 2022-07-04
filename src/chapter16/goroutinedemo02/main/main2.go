package main

import (
	"fmt"
	"sync"
	"time"
)

// 计算 1 - 200 的各个数的阶乘，并且把各个数的阶乘放入到map中
//
var (
	myMap2 = make(map[int]int, 10)
	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁
	// sync 是包：synchorinized 同步
	// Mutex: 是互斥
	lock sync.Mutex
)

func test2(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 加锁
	lock.Lock()
	myMap2[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go test2(i)
	}
	time.Sleep(time.Second)
	for k, v := range myMap2 {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
}
