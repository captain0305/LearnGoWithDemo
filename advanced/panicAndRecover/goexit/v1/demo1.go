package main

import (
	"fmt"
	"runtime"
)

func f0() int {
	var x = 1
	defer fmt.Println("正常退出：", x)
	x++
	return x
}

func f1() {
	var x = 1
	defer fmt.Println("正常退出：", x)
	x++
}

func f2() {
	var x, y = 1, 0
	defer fmt.Println("因恐慌而退出：", x)
	x = x / y // 将产生一个恐慌
	x++       // 执行不到
}

func f3() int {
	x := 1
	defer fmt.Println("因Goexit调用而退出：", x)
	x++
	runtime.Goexit()
	return x + x // 执行不到
}

func main() {
	fmt.Println(f0())
	f1()
	//f2()
	fmt.Println(f3())
}
