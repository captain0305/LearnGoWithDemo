package main

import "fmt"

// nil函数值在被调用时会产生恐慌
func panic() {
	var f func() // f == nil
	fmt.Println(f)
	f()
}

func main() {
	panic()
	defer fmt.Println("此行可以被执行到")
	var f func() // f == nil
	defer f()    // 将产生一个恐慌
	fmt.Println("此行可以被执行到")
	f = func() {} // 此行不会阻止恐慌产生
}
