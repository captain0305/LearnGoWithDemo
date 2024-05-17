package main

import (
	"fmt"
	"runtime"
	"time"
)

func f() {
	defer func() {
		fmt.Println(recover())
	}()

	// 此调用产生的Goexit信号恢复之前产生的恐慌。
	defer runtime.Goexit()
	panic("bye")

	fmt.Println("hello")
}

func main() {
	go f()

	for runtime.NumGoroutine() > 1 {
		fmt.Println(runtime.NumGoroutine())
		// 暂停1s
		//时间戳
		start := time.Now()
		//把当前goroutine从当前线程退出，放回队列
		runtime.Gosched()
		end := time.Now()
		fmt.Printf("running time is: %vs\n", end.Sub(start).Seconds())

	}

}
