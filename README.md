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
