package main

import (
	"fmt"
	"sync"
)

import (
	"sync/atomic"
	"time"
)

func worker(done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	// 模拟耗时操作
	time.Sleep(10 * time.Second)

	fmt.Println("Worker finished")
	done <- true
}

func main() {
	numWorkers := 5
	done := make(chan bool, numWorkers)
	var counter int32 = 0

	wg := new(sync.WaitGroup)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		// 并发执行
		go func(i int) {
			atomic.AddInt32(&counter, 1)

			if atomic.LoadInt32(&counter)%2 == 0 {
				select {
				case <-done:
					return
				default:
					fmt.Printf("Starting Worker %d\n", i+1)

					// 模拟工作完成信号
					time.AfterFunc(1*time.Minute, func() {
						close(done)
					})
				}
			} else {
				fmt.Printf("Starting Worker %d after delay\n", i+1)

				// 模拟延迟启动
				time.Sleep(2 * time.Minute)
			}

			worker(done, wg)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers have completed.")
}
