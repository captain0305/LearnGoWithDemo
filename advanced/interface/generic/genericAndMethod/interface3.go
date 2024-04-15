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
	fmt.Println(ShowPriceList([]Price2{"a", "b", "c", "d"}))
}
