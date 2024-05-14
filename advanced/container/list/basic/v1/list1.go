package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.List{}
	l.PushFront(1)

	printList(l)

	l.PushFront(2)
	l.InsertAfter(3, l.Front())
	printList(l)

	l.Remove(l.Front())
	printList(l)

}

func printList(list list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("--------")
}
