package main

import "fmt"

type Aboutable interface {
	About() string
}

type Book struct {
	name string
	// 更多其它字段……
}

// 实现关系是隐式的不需要implements关键字
func (book *Book) About() string {
	return "Book: " + book.name
}

type MyInt int

func (MyInt) About() string {
	return "我是一个自定义整数类型"
}

func main() {
	var aboutable Aboutable
	aboutable = &Book{"Go语言101"}
	about := aboutable.About()

	fmt.Println(about)
	fmt.Println(aboutable.About())
	aboutable = MyInt(1024)
	fmt.Println(aboutable.About())
}
