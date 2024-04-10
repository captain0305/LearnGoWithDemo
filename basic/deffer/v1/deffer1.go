package main

import "fmt"

// defer 语句会将函数推迟到外层函数返回之后执行。
// 延迟函数调用是一个后进先出队列

func deffer1() {
	defer fmt.Println("The third line.")
	defer fmt.Println("The second line.")
	fmt.Println("The first line.")
}

func deffer2() {
	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	if false {
		defer fmt.Println("not reachable")
	}
	defer func() {
		defer fmt.Println("7")
		fmt.Println("3")
		defer func() {
			fmt.Println("5")
			fmt.Println("6")
		}()
		fmt.Println("4")
	}()
	fmt.Println("2")
	return
	defer fmt.Println("not reachable")
}

func Triple(n int) (r int) {
	defer func() {
		r += n // 修改返回值为 3 * n
	}()

	return n + n // <=> r = n + n; return
}

func main() {
	deffer1()
}
