package foo // import "x.y/foo"

import "fmt"

type A struct {
	n int
}

func (a A) m() {
	fmt.Println("A", a.n)
}

type I interface {
	m()
}

// 大写为导出函数和变量，接口
func Bar(i I) {
	i.m()
}
