package main

import (
	"fmt"
	"sync"
	"time"
)

func safeMap() {
	var sm sync.Map
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		//wg.done减少计数器数量
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			sm.Store(i, i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			//协程休眠操作
			time.Sleep(5 * time.Millisecond) // Sleep for 0.5 seconds
			sm.Store(i, i+1)
		}
	}()

	wg.Wait()
	//sync.map无法获取长度
	//fmt.Println("Length of normal map:", len(sm))
	len := 0
	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		len++
		return true
	})
	fmt.Println("Length of normal map:", len)
}

// 无法运行
func unsafeMap() {
	var m = make(map[int]int)

	wg := &sync.WaitGroup{}
	//设置计数器数量为2
	wg.Add(2)

	go func() {
		//wg.done减少计数器数量
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			m[i] = i

		}

	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			m[i] = i + 1
		}
	}()

	wg.Wait()

	fmt.Println("Length of normal map:", len(m))
}

func main() {
	safeMap()

	//unsafeMap()
}
