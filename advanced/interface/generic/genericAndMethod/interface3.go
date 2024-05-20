package main

import (
	"fmt"
	"strconv"
)

type Price int

type Price2 string

func (i Price) String() string {
	return strconv.Itoa(int(i))
}

func (i Price2) String() string {
	return string(i)
}

// 波浪号表示可以是衍生类型
type ShowPrice interface {
	String() string
	~int | ~string
}

func ShowPriceList[T ShowPrice](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return

}

func main() {

	fmt.Println(ShowPriceList([]Price{1, 2, 3, 4, 5}))
	fmt.Println(ShowPriceList([]Price2{"a", "demo", "c", "d"}))
}
