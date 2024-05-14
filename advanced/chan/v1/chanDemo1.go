package main

import "fmt"

func main() {
	//通道的容量默认为0，当通道中没有数据时，接收操作会阻塞

	//报错fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int, 1)
	fmt.Println(cap(ch))
	fmt.Println(len(ch))
	x := 0
	ch <- x // a send statement
	ch <- 1
	fmt.Println(len(ch))
	x = <-ch // a receive expression in an assignment statement
	fmt.Println(x)
	//<-ch // a receive statement; result is discarded
}
