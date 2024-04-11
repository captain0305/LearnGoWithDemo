package main

// 很恶心的用法不建议使用
var aBook = struct {
	author struct { // 此字段的类型为一个匿名结构体类型
		firstName, lastName string
		gender              bool
	}
	title string
	pages int
}{
	author: struct {
		firstName, lastName string
		gender              bool
	}{
		firstName: "Mark",
		lastName:  "Twain",
	}, // 此组合字面量中的类型为一个匿名结构体类型
	title: "The Million Pound Note",
	pages: 96,
}

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	_ = m
	s := []int{1, 2, 3}
	_ = s
	x := [3]int{1, 2, 3}
	_ = x

}
