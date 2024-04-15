package main

import "fmt"

type Filter interface {
	About() string
	Process([]int) []int
}

// UniqueFilter用来删除重复的数字。
type UniqueFilter struct{}

func (a UniqueFilter) About() string {
	return "删除重复的数字"
}

func (UniqueFilter) Process(inputs []int) []int {
	outs := make([]int, 0, len(inputs))
	pusheds := make(map[int]bool)
	for _, n := range inputs {
		if !pusheds[n] {
			pusheds[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}

// MultipleFilter用来只保留某个整数的倍数数字。
type MultipleFilter int

func (mf MultipleFilter) About() string {
	return fmt.Sprintf("保留%v的倍数", mf)
}
func (mf MultipleFilter) Process(inputs []int) []int {
	var outs = make([]int, 0, len(inputs))
	for _, n := range inputs {
		if n%int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

type myFilter struct{}

func (myFilter) About() string {
	fmt.Println("自定义过滤器")
	return "自定义过滤器"
}

func (myFilter) Process(inputs []int) []int {
	fmt.Println("自定义过滤器")
	return inputs

}

// 在多态特性的帮助下，只需要一个filteAndPrint函数。
func filteAndPrint(fltr Filter, unfiltered []int) []int {
	// 在fltr参数上调用方法其实是调用fltr的动态值的方法。
	filtered := fltr.Process(unfiltered)
	fmt.Println(fltr.About()+":\n\t", filtered)
	return filtered
}

func main() {
	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("过滤之前：\n\t", numbers)

	// 三个非接口值被包裹在一个Filter切片
	// 的三个接口元素中。
	//Filter是接口类型，UniqueFilter和MultipleFilter是实现Filter接口的类型
	//为什么Unique是个结构体但是还能是接口类型
	//在你的代码中，UniqueFilter 是一个结构体类型，它实现了 Filter 接口，因为它提供了 Filter 接口所需要的所有方法，即 About 和 Process 方法。
	//所以，尽管 UniqueFilter 是一个结构体，但它也可以被视为一个 Filter 接口类型。
	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(3),
		//myFilter{},
	}

	// 每个切片元素将被赋值给类型为Filter的
	// 循环变量fltr。每个元素中的动态值也将
	// 被同时复制并被包裹在循环变量fltr中。
	for _, fltr := range filters {
		numbers = filteAndPrint(fltr, numbers)
	}
}
