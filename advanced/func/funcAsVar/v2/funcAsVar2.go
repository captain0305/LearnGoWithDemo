package main

import "fmt"

func main() {
	// 此函数返回一个函数类型的结果，亦即闭包（closure）。
	/*	函数闭包（Closure）是一种特殊的函数，它可以捕获并保存其外部作用域中的变量，
		即使这个函数在其原始作用域之外被调用。
		这意味着闭包函数可以“记住”和访问其定义时的环境。
	*/
	//第一层函数返回一个函数，第二层函数返回一个bool值
	//第一层函数闭包
	isMultipleOfX := func(x int) func(int) bool {
		return func(n int) bool {
			return n%x == 0
		}
	}
	//第二层函数闭包
	var isMultipleOf3 = isMultipleOfX(3)
	var isMultipleOf5 = isMultipleOfX(5)
	fmt.Println(isMultipleOf3(6))  // true
	fmt.Println(isMultipleOf3(8))  // false
	fmt.Println(isMultipleOf5(10)) // true
	fmt.Println(isMultipleOf5(12)) // false

	isMultipleOf15 := func(n int) bool {
		return isMultipleOf3(n) && isMultipleOf5(n)
	}
	fmt.Println(isMultipleOf15(32)) // false
	fmt.Println(isMultipleOf15(60)) // true
}
