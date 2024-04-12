package main

import "fmt"

func summarize1() {
	//不可寻址但可以取地址的值
	pm := &map[string]int{"C": 1972, "Go": 2009}
	pmm := map[string]int{"C": 1972, "Go": 2009}
	ps := &[]string{"break", "continue"}
	pa := &[...]bool{false, true, true, false}
	fmt.Println(pm, ps, pa)

	(*pm)["C"] = 1974
	pmm["d"] = 2009

	//格式化换行符
	//fmt.Printf("%T\n", pm) // *map[string]int
	//fmt.Printf("%T\n", ps) // *[]string
	//fmt.Printf("%T\n", pa) // *[4]bool

	book1 := &Book{100} // book1是一个指针
	fmt.Println(book1)

	book1.pages = 200
	a := &book1.pages
	fmt.Println(a)
}

type myMap map[string]int

type a int

type mySlice []string

func (mySlice) String() string {
	return "123"

}

func (myMap) String() string {
	return "123"
}

type Book struct {
	pages int
}

func (b Book) String() string {
	return fmt.Sprintf(" Pages: %d", b.pages)
}

func main() {
	summarize1()
}

func summarize2() {
	a := 1
	p := &a
	fmt.Println(p) // 输出：1

	m := &map[string]int{"key": 1}
	fmt.Println(m) // 输出：1

	s := &[]int{1, 2, 3}

	fmt.Println(s) // 输出：1

}

// 内嵌组合字面量可以被简化
func summarize3() {
	var heads = []*[4]byte{
		&[4]byte{'P', 'N', 'G', ' '},
		&[4]byte{'G', 'I', 'F', ' '},
		&[4]byte{'J', 'P', 'E', 'G'},
	}
	var head = []*[4]byte{
		{'P', 'N', 'G', ' '},
		{'G', 'I', 'F', ' '},
		{'J', 'P', 'E', 'G'},
	}

	fmt.Println(heads, head)

	type language struct {
		name string
		year int
	}

	var _ = [...]language{
		language{"C", 1972},
		language{"Python", 1991},
		language{"Go", 2009},
	}

	var _ = [...]language{
		{"C", 1972},
		{"Python", 1991},
		{"Go", 2009},
	}

	type LangCategory struct {
		dynamic bool
		strong  bool
	}

	// 此映射值的类型的键值类型为一个结构体类型，
	// 元素类型为另一个映射类型：map[string]int。
	var _ = map[LangCategory]map[string]int{
		LangCategory{true, true}: map[string]int{
			"Python": 1991,
			"Erlang": 1986,
		},
		LangCategory{true, false}: map[string]int{
			"JavaScript": 1995,
		},
		LangCategory{false, true}: map[string]int{
			"Go":   2009,
			"Rust": 2010,
		},
		LangCategory{false, false}: map[string]int{
			"C": 1972,
		},
	}

	var _ = map[LangCategory]map[string]int{
		{true, true}: {
			"Python": 1991,
			"Erlang": 1986,
		},
		{true, false}: {
			"JavaScript": 1995,
		},
		{false, true}: {
			"Go":   2009,
			"Rust": 2010,
		},
		{false, false}: {
			"C": 1972,
		},
	}
}

func summarize4() {
	var a [16]byte
	var s []int
	var m map[string]int

	fmt.Println(a == a)                  // true
	fmt.Println(m == nil)                // true
	fmt.Println(s == nil)                // true
	fmt.Println(nil == map[string]int{}) // false
	fmt.Println(nil == []int{})          // false

}

func summarize5() {
	m := *new(map[string]int)   // <=> var m map[string]int
	fmt.Println(m == nil)       // true
	s := *new([]int)            // <=> var s []int
	fmt.Println(s == nil)       // true
	a := *new([5]bool)          // <=> var a [5]bool
	fmt.Println(a == [5]bool{}) // true
}
