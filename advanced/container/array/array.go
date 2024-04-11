package main

import "fmt"

const Size = 32

type Person struct {
	name string
	age  int
}

func array() {
	// 数组类型
	var a [5]string
	var b [Size]int
	var c [16][]byte  // 元素类型为一个切片类型：[]byte
	var d [100]Person // 元素类型为一个结构体类型：Person
	_ = a
	_ = b
	_ = c
	_ = d
}

func array2() {
	// 下面这些数组字面量都是等价的。
	_ = [4]bool{false, true, true, false}
	_ = [4]bool{0: false, 1: true, 2: true, 3: false}
	_ = [4]bool{1: true, true}
	_ = [4]bool{2: true, 1: true}
	_ = [...]bool{false, true, true, false}
	a := [...]bool{3: false, 1: true, true}

	fmt.Println(a[0])
	fmt.Println(a)
}

func main() {
	array2()
}
