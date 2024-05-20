package main

import "fmt"

type MySlice[T int | float64] []T

func (s MySlice[T]) sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func sum() {
	var a MySlice[int] = []int{1, 2, 3, 4, 5}
	var b MySlice[float64] = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	//var c MySlice[string] = []string{"a", "demo", "c", "d", "e"}

	fmt.Println(a.sum())
	fmt.Println(b.sum())
}

func Add[T int | float64 | string](a T, b T) T {
	return a + b

}

func ADD() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.1, 2.2))
	fmt.Println(Add("a", "demo"))
}

func main() {
	ADD()
}
