package main

import "fmt"

func ergodic() {
	m := map[string]int{"C": 1972, "C++": 1983, "Go": 2009}
	for lang, year := range m {
		fmt.Printf("%v: %v \n", lang, year)
	}

	a := [...]int{2, 3, 5, 7, 11}
	for i, prime := range a {
		fmt.Printf("%v: %v \n", i, prime)
	}

	s := []string{"go", "defer", "goto", "var"}
	for i, keyword := range s {
		fmt.Printf("%v: %v \n", i, keyword)
	}
}

/*
1.被遍历容器值是原来容器的一个副本
2.如果aContainer是一个数组，那么在遍历过程中对此数组元素的修改不会体现到循环变量中。 原因是此数组的副本（被真正遍历的容器）和此数组不共享任何元素。
3.如果aContainer是一个切片（或者映射），那么在遍历过程中对此切片（或者映射）元素的修改将体现到循环变量中。

	原因是此切片（或者映射）的副本和此切片（或者映射）共享元素（或条目）。

4.如果要修改原来的值要person[i].name
*/
type Person struct {
	name string
	age  int
}

func travelsal() {

	persons := [2]Person{{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)
		// 此修改将不会体现在这个遍历过程中，
		// 因为被遍历的数组是persons的一个副本。
		persons[i].name = "Jack"
		// 此修改不会反映到persons数组中，因为p
		// 是persons数组的副本中的一个元素的副本。
		p.age = 31
	}
	fmt.Println("persons:", &persons)
}

func travelsal2() {
	// 数组改为切片
	persons := []Person{{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		p.name = "aaa"
		p.age = 31
		fmt.Println(i, p)
		// 这次，此修改将反映在此次遍历过程中。
		persons[1].name = "Jack"
		// 这个修改仍然不会体现在persons切片容器中。

	}
	fmt.Println("persons:", &persons)
}

func travelsal3() {
	m := map[int]struct{ dynamic, strong bool }{
		0: {true, false},
		1: {false, true},
		2: {false, false},
	}

	for _, v := range m {
		//m[i] = struct{ dynamic, strong bool }{true, true}
		// This following line has no effects on the map m.

		v.dynamic, v.strong = true, true
	}

	fmt.Println(m[0]) // {true false}
	fmt.Println(m[1]) // {false true}
	fmt.Println(m[2]) // {false false}
}

func travelsal4() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		delete(m, "two")
		m["four"] = 4
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println(m)
}

//复制一个切片和映射代价小，但是复制大尺寸数组代价大
//核心 for k,v的话 v是一个固定地址值的变量

func main() {

}
