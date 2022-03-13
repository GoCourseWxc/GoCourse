package main

import "fmt"

func fbn(n int) int {
	if n == 1 || n == 0 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func main() {
	fmt.Println(fbn(4))
}
