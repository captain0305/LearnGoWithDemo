package main

import "fmt"

func main() {
	//向一个nil通道发送数据或者从一个nil通道接收数据将使当前协程永久阻塞。
	var a chan int
	a <- 1
	fmt.Println(a)
}
