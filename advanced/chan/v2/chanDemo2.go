package main

import "fmt"

func main() {
	ch1 := make(chan<- int, 10) //只能发送
	ch2 := make(<-chan int, 10) //只能接收
	ch3 := make(chan int, 10)   //双向通道

	fmt.Println(ch1, ch2, ch3)
	ch1 <- 1
	ch1 <- 2

	close(ch1)

	//panic: send on closed channel
	//ch1 <- 3

	//报错只能是双向通道或发送通道才能关闭
	//close(ch2)

	close(ch3)

}
