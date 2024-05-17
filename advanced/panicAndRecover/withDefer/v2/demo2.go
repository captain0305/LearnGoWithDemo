package main

import (
	"fmt"
	"time"
)

// 捕获后不做处理
func recover1() {
	defer func() {
		fmt.Println("正常退出")
	}()
	fmt.Println("嗨！")
	defer func() {
		v := recover()
		fmt.Println("恐慌被恢复了：", v)
	}()
	panic("拜拜！") // 产生一个恐慌
	fmt.Println("执行不到这里")
}

// 恐慌捕获后做处理
func recover2() {
	fmt.Println("hi!")

	//协程内部产生恐慌不做处理，此协程会在恐慌状态下退出导致崩溃
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered from ", r)
			}

		}()
		time.Sleep(time.Second)
		panic(123)
		fmt.Println("执行不到这里")

	}()

	time.Sleep(10 * time.Second)
	fmt.Println("能执行到这里")
}

func main() {
	recover1()
	recover2()
}
