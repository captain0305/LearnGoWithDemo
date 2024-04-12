package main

import "fmt"

// 泛型类型
type Slice[T int | float64 | float32] []T

type MyMap[KEY int | string, VALUE any] map[KEY]VALUE

type MyStruct[T int | string] struct {
	Id   T
	Name string
}

// 泛型接口
type IprintData[T int | float32 | string] interface {
	PrintData(data T) string
}

type MyChan[T int | string] chan T

// 泛型函数
func Sum[T int | float64 | float32](s Slice[T]) (sum T) {
	for _, v := range s {
		sum = v + sum
	}
	return
}

func main() {
	var a Slice[int] = []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	var b Slice[float64] = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(b)
	var c Slice[float32] = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(c)

	fmt.Println(Sum(a))
	fmt.Println(Sum(b))
	fmt.Println(Sum(c))

}
