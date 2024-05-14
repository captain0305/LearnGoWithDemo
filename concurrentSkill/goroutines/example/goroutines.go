package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	wg := new(sync.WaitGroup)
	var j int64 = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			atomic.AddInt64(&j, 1)
		}(i)
	}

	wg.Wait()
	fmt.Println("result:", atomic.LoadInt64(&j))
}
