package main

import (
	"fmt"
)

// 对于有名的：第一步先执行defer，对变量进行修改；第二步，返回被修改的返回值。
func test() (i int) {
	defer func() {
		i++
	}()
	return 1
}

// 对于匿名的：第一步设置临时变量保存返回值；第二步按照defer的执行步骤执行defer语句，如果其中有对变量的修改，将不会影响s变量的值。
func test2() int {
	i := 1
	defer func() {
		i++
	}()
	return i
}

func main() {
	fmt.Println(test())
	fmt.Println(test2())

}
