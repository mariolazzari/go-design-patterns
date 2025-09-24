package main

import (
	"fmt"
)

// Object Interface
type Animal interface {
	MakeSound() string
}

type Cat struct{}

func (c *Cat) MakeSound() string {
	return "Meow"
}

type Dog struct{}

func (d *Dog) MakeSound() string {
	return "Woof"
}

// Strategy Interface
type Strategy interface {
	Execute() string
}

// Concrete Strategy A
type SpeakStrategy struct {
	animal Animal
}

func (s *SpeakStrategy) Execute() string {
	return s.animal.MakeSound()
}

// Concrete Strategy B
type BarkStrategy struct {
	animal Animal
}

func (s *BarkStrategy) Execute() string {
	return s.animal.MakeSound()
}

// Context
type Context struct {
	strategy Strategy
}

func (c *Context) ExecuteStrategy() string {
	return c.strategy.Execute()
}

func main() {
	cat := &Cat{}
	dog := &Dog{}

	speakStrategy := &SpeakStrategy{animal: cat}
	barkStrategy := &SpeakStrategy{animal: dog}

	context := &Context{strategy: speakStrategy}
	fmt.Println(context.ExecuteStrategy())

	context.strategy = barkStrategy
	fmt.Println(context.ExecuteStrategy())

}
