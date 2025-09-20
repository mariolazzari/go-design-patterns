# Design Patterns in Go for Object Oriented Programming

## Overview of Design Patterns

### What are design patterns

- Generalized solutions to recurrent problems
- Describe how objects and classes may depend or interact
- Create a formalized approach
- Solve specific problem
- Provide context
- Known use cases

### Different types of design patterns

- Creational
- Structural
- Behavioral

#### Creational pattern

- Provide interface for creating objects
- Construct complex objects
- Limit dependencies

#### Structural patterns

- Allows incompatible objects to comunicate
- Separate abstraction and implementation
- Attach behavior to objects

#### Behavioral patterns

- Pass information between objects
- Notify other objects
- Separate objects and albgorithms

### OOP in Go

- Structs to define types
- Hold state only
- Functions are behaviors
  
```go
package main

import "fmt"

type Person struct {
 name string
 age  int
}

type Details interface{
    print_details() string
}

func (p *Person) print_details() {
 fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

func main() {
 mario := Person{name: "Mario", age: 50}
 mario.print_details()
}
```

## Creational patterns

### What are creational patterns

- Manage objects creation
- Encapsulate knowledge of the classes
- Hide instances

### Classification of creational patterns

- Object creational: defer object creation to another object
- Class creational: defer class creation to subclasses

### Use cases of creational patterns

- Construct different representation of complex objects
- Hide implementation of class or object
- Allow subclasses to alter the type of the created object
- On instance only

### Builder pattern

- Construct object step by step
- Each step is independent

#### Usage of builder pattern

- Rid of telescoping constructor
- Different representation of same object
- Construct tree or complex objects

#### Builder pattern pros

- Code reuse
- Defer steps
- Isole complex code

#### Builder patter cons

- Complex code

### Implement a builder pattern

```go
package main

import (
 "fmt"
)

// Product
type Burger struct {
 breadType string
 meatType  string
 toppings  string
}

// Builder Interface
type BurgerBuilder interface {
 setBreadType()
 setMeatType()
 setToppings()
 getBurger() Burger
}

// Concrete Builder 1
type RegularBurgerBuilder struct {
 breadType string
 meatType  string
 toppings  string
}

func newRegularBurgerBuilder() *RegularBurgerBuilder {
 return &RegularBurgerBuilder{}
}

func (b *RegularBurgerBuilder) setBreadType() {
 b.breadType = "Sesame"
}

// Set the Meat Type
func (b *RegularBurgerBuilder) setMeatType() {
 b.meatType = "beef"
}

// Set the Toppings
func (b *RegularBurgerBuilder) setToppings() {
 b.toppings = "Lettuce, Tomato, Bacon, and Cheese"
}

func (b *RegularBurgerBuilder) getBurger() Burger {
 return Burger{
  breadType: b.breadType,
  meatType:  b.meatType,
  toppings:  b.toppings,
 }
}

// Concrete Builder 2
type VeganBurgerBuilder struct {
 breadType string
 meatType  string
 toppings  string
}

func newVeganBurgerBuilder() *VeganBurgerBuilder {
 return &VeganBurgerBuilder{}
}

func (b *VeganBurgerBuilder) setBreadType() {
 b.breadType = "Gluten Free"
}

// Set the Meat Type
func (b *VeganBurgerBuilder) setMeatType() {
 b.meatType = "Black Bean"
}

// Set the Toppings
func (b *VeganBurgerBuilder) setToppings() {
 b.toppings = "Lettuce and Tomato"
}

func (b *VeganBurgerBuilder) getBurger() Burger {
 return Burger{
  breadType: b.breadType,
  meatType:  b.meatType,
  toppings:  b.toppings,
 }
}

// Director
type Director struct {
 builder BurgerBuilder
}

func newDirector(b BurgerBuilder) *Director {
 return &Director{
  builder: b,
 }
}

func (d *Director) setBuilder(b BurgerBuilder) {
 d.builder = b
}

func getBuilder(builderType string) BurgerBuilder {
 if builderType == "regular" {
  return newRegularBurgerBuilder()
 }

 if builderType == "vegan" {
  return newVeganBurgerBuilder()
 }

 return nil
}

func (d *Director) buildBurger() Burger {
 d.builder.setBreadType()
 d.builder.setMeatType()
 d.builder.setToppings()
 return d.builder.getBurger()
}

func main() {
 regularBurgerBuilder := getBuilder("regular")
 veganBurgerBuilder := getBuilder("vegan")

 // Initialize the director
 director := newDirector(regularBurgerBuilder)

 // Create the regular burger
 regularBurger := director.buildBurger()

 fmt.Printf("Regular Burger Bread Type: %s\n", regularBurger.breadType)
 fmt.Printf("Regular Burger Meat Type: %s\n", regularBurger.meatType)
 fmt.Printf("Regular Burger Toppings: %s\n", regularBurger.toppings)

 // Build the vegan burger
 director.setBuilder(veganBurgerBuilder)
 veganBurger := director.buildBurger()

 fmt.Printf("Vegan Burger Bread Type: %s\n", veganBurger.breadType)
 fmt.Printf("Vegan Burger Meat Type: %s\n", veganBurger.meatType)
 fmt.Printf("Vegan Burger Toppings: %s\n", veganBurger.toppings)

}
```
