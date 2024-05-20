package v1

// 选择器遮挡和碰撞
type A struct {
	x string
}

func (A) y(int) bool {
	return false
}

type B struct {
	y bool
}

func (B) x(string) {}

type C struct {
	B
}

var v1 struct {
	A
	B
}

func f1() {
	_ = v1.B.x // error: 模棱两可的v1.x
	_ = v1.B.y // error: 模棱两可的v1.y
	_ = v1.A.x // error: 模棱两可的v1.x
	_ = v1.A.y // error: 模棱两可的v1.x
}
