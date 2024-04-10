package main

import "fmt"

func pointer1() {
	var a *int
	a = new(int)
	fmt.Println(a)
}

func pointer2() {
	p0 := new(int)  // p0指向一个int类型的零值
	fmt.Println(p0) // （打印出一个十六进制形式的地址）
	*p0 = 1
	fmt.Println(*p0) // 0

	x := *p0         // x是p0所引用的值的一个复制。
	p1, p2 := &x, &x // p1和p2中都存储着x的地址。
	// x、*p1和*p2表示着同一个int值。
	fmt.Println(p1 == p2) // true
	fmt.Println(p0 == p1) // false
	p3 := &*p0            // <=> p3 := &(*p0)
	// <=> p3 := p0
	// p3和p0中存储的地址是一样的。
	fmt.Println(p0 == p3) // true
	*p0, *p1 = 123, 789
	fmt.Println(*p2, x, *p3) // 789 789 123

	fmt.Printf("%T, %T \n", *p0, x) // int, int
	fmt.Printf("%T, %T \n", p0, p1) // *int, *int
}

func double(x *int) {
	*x += *x
	x = nil
}
func pointer3() {
	var a = 3
	double(&a)
	fmt.Println(a)
	//地址值
	p := &a
	double(&a)
	fmt.Println(a, p == nil)
}

func newInt() *int {
	a := 3
	return &a
}

// go不支持指针参与算数运算
func pointer4() {
	a := int64(5)
	p := &a

	// 下面这两行编译不通过。
	/*
		p++
		p = (&a) + 8
	*/

	*p++
	fmt.Println(*p, a)   // 6 6
	fmt.Println(p == &a) // true

	*&a++
	fmt.Println(a)
	*&*&*&a++
	fmt.Println(a)
	**&*&p++
	fmt.Println(a)
	*&*p++
	fmt.Println(*p, a) // 10 10
}

type MyInt int64
type Ta *int64
type Tb *MyInt

// 指针类型转换
func pointer5() {
	myInt := MyInt(12)

	//int64
	i := int64(23)
	ta := Ta(&i)
	fmt.Println("隐式转换", ta)

	var tbb Tb = Tb((*MyInt)((*int64)(ta)))
	fmt.Println("三层显式转换", tbb)

	var inta int64 = int64(myInt)
	fmt.Println("显式转换", inta)

	tb := Tb(&myInt)
	fmt.Println(*tb)
	fmt.Println(tb)
}

// 指针类型比较
func pointer6() {
	// 4个不同类型的指针：
	var pa0 Ta
	var pa1 *int64
	var pb0 Tb
	var pb1 *MyInt

	// 下面这6行编译没问题。它们的比较结果都为true。
	_ = pa0 == (*int64)(pb1)
	_ = pa0 == pa1
	_ = pb0 == pb1
	_ = pa0 == nil
	_ = pa1 == nil
	_ = pb0 == nil
	_ = pb1 == nil

	// 下面这三行编译不通过。

	//_ = pa0 == pb0
	//_ = pa1 == pb1
	//_ = pa0 == Tb(nil)

}

func main() {
	pointer6()
}
