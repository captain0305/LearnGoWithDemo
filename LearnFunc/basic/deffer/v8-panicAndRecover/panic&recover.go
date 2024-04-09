package main

import "fmt"

// 模拟长程跳转
func main() {
	n := func() (result int) {
		defer func() {
			if v := recover(); v != nil {
				if n, ok := v.(int); ok {
					result = n
				}
			}
		}()

		func() {
			func() {
				func() {
					// ...
					panic(123) // 用恐慌来表示成功返回
				}()
				// ...
			}()
		}()
		// ...
		return 0
	}()
	fmt.Println(n) // 123
}
