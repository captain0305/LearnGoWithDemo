package main

import "fmt"

func for1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	for i < 20 {
		fmt.Println(i)
		i++
	}

}

func for2() {
	for i := 0; ; i++ { // 等价于：for i := 0; true; i++ {
		if i >= 10 {
			break
		}
		fmt.Println(i)
	}

	// 下面这几个循环是等价的。
	for true {
		break
	}
	for true {
		break
	}
	for {
		break
	}
	for {
		break
	}

}

func for3() {

	for i := 0; i < 3; i++ {
		fmt.Print(i)
		i := i // 这里声明的变量i遮挡了上面声明的i。
		// 右边的i为上面声明的循环变量i。
		i = 10 // 新声明的i被更改了。
		_ = i
	}

}

func for4() {
	//等价
	for i := range 12 {
		fmt.Println(i)
	}
	fmt.Println("---------")
	for i := 0; i < 12; i++ {
		fmt.Println(i)
	}
}
func for5() {
	var i = 2
	for i = 3; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------")

	for i = range 12 {
		fmt.Println(i)
	}
}

func main() {
	for1()
	for2()
	for3()
	for4()
	for5()

}
