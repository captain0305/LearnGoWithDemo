package main

import (
	"github.com/spf13/cast"
	"math/rand"
	"sync"
)

// mediaInterfaceServiceImpl 框架生成协议对象
type mediaInterfaceServiceImpl struct {
	a string
	b *mediaInterfaceServiceImpl2
}

type mediaInterfaceServiceImpl2 struct {
	a string
}

func main() {
	wg := new(sync.WaitGroup)
	defer wg.Done()
	wg.Add(1)
	m := new(mediaInterfaceServiceImpl)
	m.a = "a"
	m.b = new(mediaInterfaceServiceImpl2)
	m.b.a = "bstring"
	for i := 0; i < 10; i++ {

		go func() {
			for true {
				println(m.a)

				println(m.b.a)
				m.b.a = cast.ToString(rand.Int31())
			}
		}()
	}

	wg.Wait()

}
