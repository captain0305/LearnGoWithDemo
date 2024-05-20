package main

import "fmt"

func estimation1() {
	var f = func() {
		fmt.Println(false)
	}

	defer f()

	f = func() {
		fmt.Println(true)
	}
	defer f()
}

func estimation2() {
	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i+x)
		}
		x = 10
	}()
	fmt.Println()
	func() {
		var x = 0
		//一个匿名函数体内的表达式是在此函数被执行的时候才会被逐渐估值的，不管此函数是被普通调用还是延迟/协程调用。
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("demo:", i+x)
			}()
		}
		x = 10
	}()
}

// defer的执行顺序是栈的形式，先进后出,估值时刻在被推入延迟调用队列时估值
func main() {
	estimation2()
}
