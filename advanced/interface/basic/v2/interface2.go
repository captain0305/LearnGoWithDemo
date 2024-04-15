package main

import (
	"fmt"
	"strconv"
)

type Aboutable interface {
	About() string
	int | Book | MyInt
}

type Book struct {
	name string
	// 更多其它字段……
}

// 实现关系是隐式的不需要implements关键字
func (book Book) About() string {
	return "Book: " + book.name
}

type MyInt int

func (myint MyInt) About() string {
	return "我是一个自定义整数类型" + strconv.Itoa(int(myint))
}

func AboutMain[T Aboutable](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.About())
	}
	return

}

func main() {
	//	var aboutable Aboutable
	//	aboutable = &Book{"Go语言101"}
	//	about := aboutable.About()
	//
	//	fmt.Println(about)
	//	fmt.Println(aboutable.About())
	//	aboutable = MyInt(1024)
	//	fmt.Println(aboutable.About())
	//}
	fmt.Println(AboutMain([]Book{Book{"Go语言101"}, Book{"Go语言102"}}))
	fmt.Println(AboutMain([]MyInt{1024, 2048}))

}
