package main

import "fmt"

func outer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	closure := outer()
	fmt.Println(closure()) // 输出：1
	fmt.Println(closure()) // 输出：2
	fmt.Println(closure()) // 输出：3
}
