package a

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 外部可见
func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

// 外部可见
func (p *Person) SayHello() {
	fmt.Printf("Hello, my name is %s, I'm %d years old.\n", p.Name, p.Age)
}

// 外部不可见
func (a *Person) printPerson(p *Person) {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
