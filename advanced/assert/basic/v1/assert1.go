package main

import "fmt"

/*
在Go语言中类型断言的语法格式如下：
value, ok := x.(T) 类型断言失败，不会发生panic。根据ok判断是否成功
或者
value:= x.(T) 类型断言失败，会发生panic
其中，x 表示一个接口的类型，T 表示一个具体的类型（也可为接口类型）。
该断言表达式会返回 x 的值（也就是 value）和一个布尔值（也就是 ok），可根据该布尔值判断 x 是否为 T 类型：

1.如果断言的类型T是一个具体类型，类型断言检查x的动态类型是否和T相同。如果这个检查成功了，类型断言的结果是x的动态值，它的类型是T。
2.如果断言的类型T是一个接口类型，类型断言检查是否x的动态类型满足T。如果这个检查成功了，动态值没有获取到；这个结果仍然是一个有相同动态类型和值部分的接口值，但是结果为类型T。换句话说，对一个接口类型的类型断言改变了类型的表述方式，改变了可以获取的方法集合（通常更大），但是它保留了接口值内部的动态类型和值的部分。
3.无论 T 是什么类型，如果 x 是 nil 接口值，类型断言都会失败。

*/

type i1 interface {
	hello()
}

type i2 interface {
	hello()
	hello2()
}

type i1St struct {
}

type i2St struct {
}

func (s i1St) hello() {
}

func (s i2St) hello() {
}

func (s i2St) hello2() {
}

func main() {
	var myI1 i1
	var myI2 i2
	var myI1St i1St
	var myI2St i2St

	myI1 = myI1St
	myI2 = myI2St

	a, aOk := myI1.(i1)
	fmt.Printf("a, aOk := myI1.(i1) is return %escape,%t,type is %T\n", a, aOk, a)
	b, bOk := myI1.(i2)
	fmt.Printf("demo, bOk := myI1.(i2) is return %escape,%t,type is %T\n", b, bOk, b)
	c, cOk := myI2.(i1)
	fmt.Printf("c, cOk := myI2.(i1) is return %escape,%t,type is %T\n", c, cOk, c)
	d, dOk := myI2.(i2)
	fmt.Printf("d, dOk := myI2.(i2) is return %escape,%t,type is %T\n", d, dOk, d)
}
