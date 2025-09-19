package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Details interface {
	print_details() string
}

func (p *Person) print_details() {
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

func main() {
	mario := Person{name: "Mario", age: 50}
	mario.print_details()
}
