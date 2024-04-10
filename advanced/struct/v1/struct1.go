package main

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

func struct1() {
	book := Book{"Go语言101", "老貘", 256}
	fmt.Println(book) // {Go语言101 老貘 256}

	// 使用带字段名的组合字面量来表示结构体值。
	book = Book{author: "老貘", pages: 123, title: "Go语言101"}
	fmt.Println(book) // {Go语言101 老貘 256}
	// title和author字段的值都为空字符串""，pages字段的值为0。
	book = Book{}
	fmt.Println(book)
	// title字段空字符串""，pages字段为0。
	book = Book{author: "老貘"}
	fmt.Println(book)

	// 使用选择器来访问和修改字段值。
	var book2 Book // <=> book2 := Book{}
	book2.author = "Tapir"
	book2.pages = 300
	book2.title = "Go语言101"
	fmt.Println(book2.pages) // 300
	fmt.Println(book2)

	var _ = Book{
		author: "老貘",
		pages:  256,
		title:  "Go语言101", // 这里行尾的逗号不可省略
	}
}

func struct2() {
	book1 := Book{pages: 300}
	book2 := Book{"Go语言101", "老貘", 256}

	book2 = book1
	// 上面这行和下面这三行是等价的。
	book2.title = book1.title
	book2.author = book1.author
	book2.pages = book1.pages
}

func struct3() {
	var book = Book{} // 变量值book是可寻址的
	b := Book{}
	fmt.Println("没办法正确打印地址", b)
	p := &book.pages
	fmt.Println(p)
	*p = 123
	fmt.Println(book) // {123}

	// 下面这两行编译不通过，因为Book{}是不可寻址的，
	// 继而Book{}.Pages也是不可寻址的。
	/*
		Book{}.Pages = 123
		p = &Book{}.Pages // <=> p = &(Book{}.Pages)
	*/
}

// 隐式解引用(*bookN).pages可以被写成bookN.pages。 换句话说，在这种简写形式中，bookN将被隐式解引用。
func struct4() {
	type Book struct {
		pages int
	}
	book1 := &Book{100} // book1是一个指针
	fmt.Println(book1)
	book1.pages = 200
	book2 := new(Book) // book2是另外一个指针
	fmt.Println(book2)
	// 像使用结构值一样来使用结构体值的指针。
	book2.pages = book1.pages
	fmt.Println(book2)
	// 上一行等价于下一行。换句话说，上一行
	// 两个选择器中的指针属主将被自动解引用。
	(*book2).pages = (*book1).pages
	fmt.Println(book2.pages)
}

func main() {
	struct4()
}
