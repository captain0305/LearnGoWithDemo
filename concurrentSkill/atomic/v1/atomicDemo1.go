package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var count int64 = 0

// Increment increments the value stored in `count`.
func Increment() {
	atomic.AddInt64(&count, 1)
}

// Get returns the current value of `count`.
func Get() int64 {
	return atomic.LoadInt64(&count)
}

func main() {
	fmt.Println("Initial count:", Get())

	go func() {
		for {
			Increment()
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			Increment()
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			Increment()
			time.Sleep(3 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Final count:", Get())
}
