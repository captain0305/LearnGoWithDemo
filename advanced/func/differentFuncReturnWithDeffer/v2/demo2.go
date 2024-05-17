package main

import "fmt"

func test1() (x int) {
	defer fmt.Printf("in defer: x = %d\n", x)
	x = 7
	return 9
}

// defer执行是在x=7之后，所以x的值是7，并且传递给Printf，所以结果是：in defer: x = 7
func test2() (x int) {
	x = 7
	defer fmt.Printf("in defer: x = %d\n", x)
	return 9
}

func test3() (x int) {
	defer func() {
		// 匿名函数能访问外部函数的变量，受后续return x的结果影响
		// 也就是说【defer匿名函数内要访问的变量 x】最终会被【defer匿名函数外最终要 return 返回出去的变量 x】所影响
		fmt.Printf("in defer: x = %d\n", x)
	}()

	x = 7
	return 9
}

func test4() (x int) {
	defer func(n int) {
		// 针对【defer匿名函数内要访问的变量 n】，其值取决于在一开始遇到 defer 时入参 n 的值，
		// 也就是最开始还没被改变时的变量 x 的值，将其赋值给入参 n，起初是0
		// 这个变量 n 的值是独立的，不会受后续 return x 结果的影响。
		fmt.Printf("in defer x as parameter: x = %d\n", n)
		fmt.Printf("in defer x after return: x = %d\n", x)
	}(x)

	x = 7
	return 9
}

func main() {
	fmt.Println("test1")
	fmt.Printf("in main: x = %d\n", test1())
	fmt.Println("test2")
	fmt.Printf("in main: x = %d\n", test2())
	fmt.Println("test3")
	fmt.Printf("in main: x = %d\n", test3())
	fmt.Println("test4")
	fmt.Printf("in main: x = %d\n", test4())
}
