package main

type x string

func (x) M() {}

type y struct {
	z byte
}

type A struct {
	x
}

func (A) y(int) bool {
	return false
}

type B struct {
	y
}

func (B) x(string) {}

func main() {
	var v struct {
		A
		B
	}
	//_ = escape.x // error: 模棱两可的v.x
	//_ = escape.y // error: 模棱两可的v.y
	_ = v.M // ok. <=> escape.A.x.M
	_ = v.z // ok. <=> escape.B.y.z
}
