package main

import "fmt"

// 无法正常处理恐慌
func example1() {
	defer func() {

		func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
	}()
	panic("hello")
}

// 正常处理恐慌
func example2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("hello")
}

func example3() {

	defer func() {
		//recover() // 将恢复恐慌"byte"
		fmt.Println(recover())
	}()

	panic("bye")

}
func main() {
	//example1()
	//example2()
	example3()
}
