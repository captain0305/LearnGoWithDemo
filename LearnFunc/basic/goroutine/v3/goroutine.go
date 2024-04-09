package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
	一个处于阻塞状态的协程不会自发结束阻塞状态，它必须被另外一个协程通过某种并发同步方法来被动地结束阻塞状态。

如果一个运行中的程序当前所有的协程都出于阻塞状态，则这些协程将永远阻塞下去，程序将被视为死锁了。
当一个程序死锁后，官方标准编译器的处理是让这个程序崩溃。
*/
func main() {
	wg.Add(2)
	go func() {
		time.Sleep(time.Second * 2)
		wg.Wait() // 阻塞在此
	}()
	go func() {
		wg.Done()
	}()
	wg.Wait() // 阻塞在此
}
