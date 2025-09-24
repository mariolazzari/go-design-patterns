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

#### Creational patterns

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

## Creational design patterns

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

### Factory method pattern

- Defines an interface or an abstract class for creating object
- Subclasses create instances

#### Factory pattern usage

- Dependencies unknown
- Customize libraries
- Reuse existing objects

#### Factory patterns pros

- Loose coupling
- Class creation in a single place
- No breaking changes

#### Factory patterns cons

- More complex code

### Implement factory pattern

```go
package main

import "fmt"

// Product interface
type ICar interface {
 setCarType(carType string)
 setYear(year int)
 printDetails()
}

// Abstract Product
type Car struct {
 carType string
 year    int
}

func (c *Car) setCarType(carType string) {
 c.carType = carType
}

func (c *Car) setYear(year int) {
 c.year = year
}

func (c *Car) printDetails() {
 fmt.Printf("Car Type: %s\n", c.carType)
 fmt.Printf("Car Year: %d\n", c.year)
}

// Concrete Product 1

type LuxuryCar struct {
 Car
}

func newLuxuryCar() ICar {
 return &LuxuryCar{
  Car: Car{
   carType: "Luxury",
   year:    2023,
  },
 }
}

// Concrete Product 2
type HybridCar struct {
 Car
}

func newHybridCar() ICar {
 return &HybridCar{
  Car: Car{
   carType: "Hybrid",
   year:    2021,
  },
 }
}

// factory method
func getCar(carType string) ICar {
 if carType == "luxury" {
  return newLuxuryCar()
 }

 if carType == "hybrid" {
  return newHybridCar()
 }

 return nil
}

// main
func main() {
 luxuryCar := getCar("luxury")
 hybridCar := getCar("hybrid")

 luxuryCar.printDetails()
 hybridCar.printDetails()
}
```

### Abstract factory method

- Set of factory methods
- Abstract class or inteface create family of related objects
- Subclasses creates classes

#### Difference between factory and abstract factory

- Factory hide construction of 1 item
- Factory uses *encapsulation* to delegate
- Abstarct hide construction of families of item
- Abstrunct uses *composition* to delegate
  
#### Abstract factory usage

- Unknwn dependencies of families of item
- Future expansion

#### Abstract factory pros

- Ensure all classes to work togheter
- Loosly coupling
- Class creation in a single place
- No braking changes

#### Abstract factory cons

- More complex

### Abstract factory implementation

```go
package main

import "fmt"

// Abstract Product  1
type ICar interface {
 setCarType(carType string)
 printDetails()
}

type Car struct {
 carType string
}

func (c *Car) setCarType(carType string) {
 c.carType = carType
}

func (c *Car) printDetails() {
 fmt.Printf("Car Type: %s\n", c.carType)
}

// Abstract Product  2
type ICarDetails interface {
 setTransmission(transmission string)
 setEngine(engine string)
 setGasType(gasType string)
 printDetails()
}

type CarDetails struct {
 transmission string
 engine       string
 gasType      string
}

func (cd *CarDetails) setTransmission(transmission string) {
 cd.transmission = transmission
}

func (cd *CarDetails) setEngine(engine string) {
 cd.engine = engine
}

func (cd *CarDetails) setGasType(gasType string) {
 cd.gasType = gasType
}

func (cd *CarDetails) printDetails() {
 fmt.Printf("Car Transmission: %s\n", cd.transmission)
 fmt.Printf("Car Engine: %s\n", cd.engine)
 fmt.Printf("Car Gas Type: %s\n", cd.gasType)
 fmt.Println()
}

// Concrete Product A1
type LuxuryCar struct {
 Car
}

// Concrete Product A2
type LuxuryCarDetails struct {
 CarDetails
}

// Concrete Product B1
type HybridCar struct {
 Car
}

// Concerete Product B2
type HybridCarDetails struct {
 CarDetails
}

// Abstract Factory Interface
type ICarFactory interface {
 makeCar() ICar
 makeCarDetails() ICarDetails
}

// Concerete Factoy 1

type LuxuryCarFactory struct {
}

func (l *LuxuryCarFactory) makeCar() ICar {
 return &LuxuryCar{
  Car: Car{
   carType: "Luxury",
  },
 }
}

func (l *LuxuryCarFactory) makeCarDetails() ICarDetails {
 return &LuxuryCarDetails{
  CarDetails: CarDetails{
   transmission: "manual",
   engine:       "gas",
   gasType:      "premium",
  },
 }
}

type HybridCarFactory struct {
}

func (l *HybridCarFactory) makeCar() ICar {
 return &HybridCar{
  Car: Car{
   carType: "Hybrid",
  },
 }
}

func (l *HybridCarFactory) makeCarDetails() ICarDetails {
 return &HybridCarDetails{
  CarDetails: CarDetails{
   transmission: "automatic",
   engine:       "hybrid",
   gasType:      "electric",
  },
 }
}

// factory method
func getCarFactory(carType string) ICarFactory {
 if carType == "luxury" {
  return &LuxuryCarFactory{}
 }

 if carType == "hybrid" {
  return &HybridCarFactory{}
 }

 return nil
}

func main() {

 luxuryFactory := getCarFactory("luxury")
 hybridFactory := getCarFactory("hybrid")

 // Build the luxury car
 luxuryCar := luxuryFactory.makeCar()
 luxuryCarDetails := luxuryFactory.makeCarDetails()

 luxuryCar.printDetails()
 luxuryCarDetails.printDetails()

 hybridCar := hybridFactory.makeCar()
 hybridCarDetails := hybridFactory.makeCarDetails()

 hybridCar.printDetails()
 hybridCarDetails.printDetails()
}
```

### Prototype pattern

- Creates copy of an object without being costly
- Object is not dependent to class code

#### Prototype pattern usage

- Code is not dependent on specific class
- Reduce amount of subclasses
- Classes initialized at runtime

#### Prototype pattern pros

- Two different objects not coupled
- Remove repeated initialization
- Avoid inheritance

#### Prototype pattern cons

- Complex objects are difficult to copy

### Prototype pattern implementation

```go
package main

import "fmt"

// Prototype interface
type Shape interface {
 print()
 clone() Shape
}

// Concrete prototype
type Circle struct {
 color string
}

func (c *Circle) print() {
 fmt.Println(c.color + " circle")

}

func (c *Circle) clone() Shape {
 fmt.Println("Cloning...")
 return &Circle{color: c.color}
}

func main() {

 greenCircle := &Circle{color: "Green"}
 blueCircle := &Circle{color: "Blue"}
 redCircle := &Circle{color: "Red"}

 greenClone := greenCircle.clone()
 blueClone := blueCircle.clone()
 redClone := redCircle.clone()

 greenCircle.print()
 blueCircle.print()
 redCircle.print()
 greenClone.print()
 blueClone.print()
 redClone.print()

}
```

### Singleton pattern

- Single class creates an object
- Class make sure only 1 object created

#### Singleton pattern usages

- Single instance
- Global access point

#### Singleton pattern pros

- 1 instance only
- Global access
- 1 initialization first time

#### Singleton pattern cons

- Against signle responsability principle
- Difficult to test
- Special handling with multiple threads

#### Singleton pattern implementation

```go
package main

import (
 "fmt"
 "sync"
)

var lock = &sync.Mutex{}

type Database struct {
}

var database *Database

func getInstance() *Database {
 if database == nil {
  lock.Lock()
  defer lock.Unlock()
  if database == nil {
   fmt.Println("Creating a single database")
   database = &Database{}
  } else {
   fmt.Println("Database has already been created")
  }
 } else {
  fmt.Println("Database has already been created")
 }
 return database

}

func main() {
 for range 5 {
  getInstance()
 }
}
```

## Structural design patterns

### What are structural patterns?

- Descrobe how classes and objects are combined
- Simplify structure with relations
- More extensibility and flexibility

### Structural patterns use cases

- Allow incompatible interfaces to collaborate
- Compese objects in tree strcture
- Attach new behaviors to objects
- Create simple interfaces

### Adapter pattern

- Bridge between 2 incopatible objects
- Single itnerface to join features
- Both sides are unaware about adapted

#### Adapter pattern usage

- Use a specific class incompatible to the rest of code
- Reuse subclasses with missing features
  
#### Adapter pattern types

- Class
- Object

#### Adapter pattern pros

- Single responsibility principle
- Open-closed principle

#### Adapter pattern cons

- Complex due lots of interfaces and classes

### Adapter pattern implementation

```go
package main 
import "fmt"

// Client
type Publisher struct {
}

func (p *Publisher) publishContentOnPlatform(platform Platform) {
    fmt.Println("Publisher is ready to publish your content.")
    platform.postMedia()
}

//Client Interface
type Platform interface {
    postMedia()
}

// Compatible Service
type Instagram struct {
}

func (i *Instagram) postMedia() {
    fmt.Println("Instagram has published your post.")
}

// Incompatible Service
type TikTok struct {
}

func (t *TikTok) scheduleMedia() {
    fmt.Println("TikTok is ready to schedule your post.")
}

// Adapter
type TikTokAdapter struct {
    tikTok *TikTok
}

func (t *TikTokAdapter) postMedia() {
   t.tikTok.scheduleMedia()
   fmt.Println("Adapter has posted the TikTok content.")
}

func main() {

    publisher := &Publisher{}
    instagram := &Instagram{}

    publisher.publishContentOnPlatform(instagram)

    tikTok := &TikTok{}
    tikTokAdapter := &TikTokAdapter{
        tikTok: tikTok,
    }

    publisher.publishContentOnPlatform(tikTokAdapter)
}
```

### Composite pattern

- Tree like structure
- Simple and complex objects treated the same way

#### Composite pattern pros

- Tree structure is convenient
- Open-closed principle

#### Composite pattern cons

- Common interface is difficult to provide

### Composite pattern implementation

```go
package main

import "fmt"

// Component interface
type Member interface {
 printMemberInfo()
}

// Leaf
type TeamMember struct {
 name       string
 teamNumber int
 position   string
}

func (t *TeamMember) printMemberInfo() {
 fmt.Printf("Name: %s Team Number: %d Position: %s\n", t.name, t.teamNumber, t.position)
}

// Composite
type Roster struct {
 members []Member
 name    string
}

// We'll implement the printMemberInfo method
func (r *Roster) printMemberInfo() {
 fmt.Println("Here's the roster for team: " + r.name)
 for _, member := range r.members {
  member.printMemberInfo()
 }
}

func (r *Roster) add(m Member) {
 r.members = append(r.members, m)
}

func main() {

 member1 := &TeamMember{name: "Johnny Rocket", teamNumber: 12, position: "Forward"}
 member2 := &TeamMember{name: "Tim Hoops", teamNumber: 24, position: "Point Guard"}
 member3 := &TeamMember{name: "Billy Banks", teamNumber: 29, position: "Shooting Guard"}

 roster := &Roster{
  name: "Bobcats",
 }

 roster.add(member1)
 roster.add(member2)
 roster.add(member3)

 roster.printMemberInfo()
}
```

### Decorator pattern

- Dinamically add features to an object
- Wrap original object
- Does no affect other objetcs

#### Decorator Pattern usage

- Assign additional functionalities
- Difficult extension with inheritance

#### Decorator pattern pros

- More flexible than subclasses
- Easier to add and remove repsonsability
- Single responsibility principle

#### Decorator pattern cons

- Difficult setup
- Code dependency

#### Decorator pattern implementation

```go
package main

import "fmt"

// Component Interface
type IceCreamInterface interface {
 addToppings() string
}

// Concrete Component
type IceCream struct {
 flavor string
}

func (i *IceCream) addToppings() string {
 return "2 Scoops of " + i.flavor
}

// Base Decorater
type IceCreamDecorator interface {
 makeIceCream()
}

// Concrete Decorater 1
type Sprinkles struct {
 iceCream IceCreamInterface
}

func (s *Sprinkles) addToppings() string {
 currentOrder := s.iceCream.addToppings()
 return currentOrder + " and Rainbow Sprinkles"
}

func (s *Sprinkles) makeIceCream() {
 fmt.Println("Here's Your Ice Cream Order")
 fmt.Println(s.addToppings())
 fmt.Println()
}

// Concrete Decorator 2
type Syrup struct {
 iceCream IceCreamInterface
}

func (s *Syrup) addToppings() string {
 currentOrder := s.iceCream.addToppings()
 return currentOrder + " and Chocolate Syrup"
}

func (s *Syrup) makeIceCream() {
 fmt.Println("Here's Your Ice Cream Order")
 fmt.Println(s.addToppings())
 fmt.Println()
}

func main() {

 iceCream := &IceCream{flavor: "Chocolate"}

 iceCreamWithSprinkles := &Sprinkles{iceCream: iceCream}

 iceCreamWithSprinkles.makeIceCream()

 iceCreamWithSprinklesAndChocolateSyrup := &Syrup{iceCream: iceCreamWithSprinkles}
 iceCreamWithSprinklesAndChocolateSyrup.makeIceCream()

}
```

### Facade pattern

- Hide complexity
- Single class
- More generic features

#### Facade patterns usage

- Simple itnerface to complex system
- Layers

#### Facade patterns pros

- Isolate complexity

#### Facade patterns cons

- Coupled to parts

#### Facade patterns implementation

```go
package main

import "fmt"

// Subsystem Class
type Post interface {
 publish()
}

type InstagramPost struct {
}

func (i *InstagramPost) publish() {
 fmt.Println("Publishing your post to Instagram")
}

type TikTokPost struct {
}

func (t *TikTokPost) publish() {
 fmt.Println("Publishing your video to TikTok")
}

type TwitterPost struct {
}

func (t *TwitterPost) publish() {
 fmt.Println("Publishing your tweet to Twitter")
}

// Facade
type Publisher struct {
 instagram *InstagramPost
 tikTok    *TikTokPost
 twitter   *TwitterPost
}

func newPublisher() *Publisher {
 fmt.Println("Initializing Publisher...")
 publisher := &Publisher{
  instagram: &InstagramPost{},
  tikTok:    &TikTokPost{},
  twitter:   &TwitterPost{},
 }
 fmt.Println("...Publisher is ready")
 return publisher
}

func (p *Publisher) postToInstagram() {
 p.instagram.publish()
}

func (p *Publisher) postToTikTok() {
 p.tikTok.publish()
}

func (p *Publisher) postToTwitter() {
 p.twitter.publish()
}

func main() {
 publisher := newPublisher()

 publisher.postToInstagram()
 publisher.postToTikTok()
 publisher.postToTwitter()
}
```

## Behavioral design patterns

### What are behavioral patterns

- Focus on algorithms
- Common communication patterns

#### Behavioral patterns usage

- Traversing collection elements
- Reducing dependencies
- Altering behaviors
- Groups of algorithms

### Iterator pattern

- Traerse elements of a collection
- Hide collection implementation

### Iterator pattern usage

- Hide complexity
- Reduce duplicate code
- Traverse unknown collection structure

### Iterator pattern pros

- Single responsbility
- Open-closed principle

### Iterator pattern cons

- Simple collecitons only
- Leff efficent

### Iterator pattern implementation

```go
package main

import (
 "fmt"
)

// Iterator Interface
type Iterator interface {
 hasNext() bool
 next() interface{}
}

// Collection Interface
type Collection interface {
 getIterator() Iterator
}

type IceCreamFlavor struct {
 name string
}

// Concrete Iterator
type IceCreamIterator struct {
 flavors []IceCreamFlavor
 current int
}

func (i *IceCreamIterator) hasNext() bool {
 return i.current < len(i.flavors)
}

func (i *IceCreamIterator) next() interface{} {

 flavor := i.flavors[i.current]
 i.current++
 return flavor
}

// Concrete Collection
type IceCreamShop struct {
 flavors []IceCreamFlavor
}

func (s *IceCreamShop) getIterator() Iterator {
 return &IceCreamIterator{
  flavors: s.flavors,
 }
}

func main() {
 shop := &IceCreamShop{
  flavors: []IceCreamFlavor{
   {"Chocolate"},
   {"Vanilla"},
   {"Pistachio"},
   {"Cookies & Cream"},
  },
 }

 iterator := shop.getIterator()
 for iterator.hasNext() {
  flavor := iterator.next().(IceCreamFlavor)
  fmt.Printf("%s\n", flavor.name)
 }
}
```

### State pattern

- State machine
- Behavior depending on object state
- Object appears to cheange its class

#### State pattern usage

- Different behavior depending on state
- Lots of conditions in class
- Duplicate code

#### State pattern pros

- Single responsibility
- Open-closed principle

#### State pattern cons

- Too much for fewer states

#### State pattern implementation

```go
package main

import (
 "fmt"
)

// State Interface
type ShapeState interface {
 drawShape()
 eraseShape()
}

// Concrete States
type CircleState struct{}

func (c *CircleState) drawShape() {
 fmt.Println("Drawing Circle")
}

func (c *CircleState) eraseShape() {
 fmt.Println("Erasing Circle")
}

type RectangleState struct{}

func (r *RectangleState) drawShape() {
 fmt.Println("Drawing Rectangle")
}

func (r *RectangleState) eraseShape() {
 fmt.Println("Erasing Rectangle")
}

// Context
type Shape struct {
 state ShapeState
}

func (s *Shape) setState(state ShapeState) {
 s.state = state
}

func (s *Shape) getState() ShapeState {
 return s.state
}

func (s *Shape) draw() {
 s.state.drawShape()
}

func (s *Shape) erase() {
 s.state.eraseShape()
}

func main() {
 circle := &Shape{}
 circle.setState(&CircleState{})
 circle.draw()
 circle.erase()

 rectangle := &Shape{}
 rectangle.setState(&RectangleState{})
 rectangle.draw()
 rectangle.erase()

}
```

### Template method pattern

- Skeleton on an algorithm in abstract class
- Subclasses override methods without changing the structure

#### Template method pattern usage

- Extends only particular parts of an algorithm
- Several classes with almost identical algorithms

#### Template method pattern pros

- Easy
- For frameworks
- Avoid duplication

#### Template method pattern cons

- Hard to mantain
- Client side only

#### Template method pattern implementation

```go
package main

import (
 "fmt"
)

// Interface
type ICar interface {
 StartEngine()
 StopEngine()
 Drive()
}

// Abstract Class
type Car struct {
 icar ICar
}

func (ac *Car) StartEngine() {
 fmt.Println("Starting engine")
}

func (ac *Car) StopEngine() {
 fmt.Println("Stopping engine")
}

// Template Method
func (c *Car) Run() {
 c.icar.StartEngine()
 c.icar.Drive()
 c.icar.StopEngine()
}

type Sedan struct {
 Car
}

func (s *Sedan) Drive() {
 fmt.Println("Driving a sedan")
}

type SUV struct {
 Car
}

func (s *SUV) Drive() {
 fmt.Println("Driving an SUV")
}

func main() {

 sedan := Car{&Sedan{}}
 sedan.Run()
 fmt.Println()
 suv := Car{&SUV{}}
 suv.Run()
}
```

### Command design pattern

- Object encapsulates request information
- Performs on trigger event

#### Command design pattern usages

- Pass method as parameter
- Schedule events
- Reversible operations

#### Command design pattern pros

- Single responsibility
- Open-closed
- Fexibility

#### Command design pattern cons

- More complicated


#### Command design pattern implementation

```go
package main

import (
 "fmt"
)

type Train struct {
 name       string
 location   string
 passengers []string
}

// Command Interface
type Command interface {
 Execute(*Train)
}

// Concrete Command 1
type AddPassengerCommand struct {
 passenger string
}

func (c *AddPassengerCommand) Execute(train *Train) {
 train.passengers = append(train.passengers, c.passenger)
 fmt.Println("New passenger on board: " + c.passenger)
}

// Concrete Command 2
type MoveTrainCommand struct {
 location string
}

func (c *MoveTrainCommand) Execute(train *Train) {
 train.location = c.location
 fmt.Println("The train is located at: " + train.location)
}

// Invoker
type Invoker struct {
 command Command
}

func (i *Invoker) ExecuteCommand(train *Train) {
 i.command.Execute(train)
}

func main() {
 train := &Train{
  name:       "Express",
  location:   "Location A",
  passengers: []string{},
 }

 addPassangerCommand := &AddPassengerCommand{passenger: "Alice"}
 moveTrainCommand := &MoveTrainCommand{location: "Station B"}

 invoker := &Invoker{}
 invoker.command = addPassangerCommand
 invoker.ExecuteCommand(train)

 invoker.command = moveTrainCommand
 invoker.ExecuteCommand(train)
}
```
