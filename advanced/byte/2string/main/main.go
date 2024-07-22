package main

import (
	"LearnGoWithDemo/advanced/byte/2string/escape"
	"fmt"
	"github.com/spf13/cast"
	"math/rand/v2"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	for i := 0; i < 10; i++ {
		go func() {
			for true {
				toString := cast.ToString(rand.Float64())
				fmt.Println("goroutine ", i, "num", toString)
				fmt.Println("goroutine ", i, "result", escape.EscapeString(toString))
			}
		}()
	}

	wg.Wait()
}
