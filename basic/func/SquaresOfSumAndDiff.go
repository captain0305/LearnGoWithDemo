package _func

func SquaresOfSumAndDiff(a int64, b int64) (s int64, d int64) {
	x, y := a+b, a-b
	s = x * x
	d = y * y
	return // <=> return s, d
}

// SquaresOfSumAndDiff2 和 SquaresOfSumAndDiff 功能一样，只是写法不同
func SquaresOfSumAndDiff2(a int64, b int64) (int64, int64) {
	return (a + b) * (a + b), (a - b) * (a - b)
}
