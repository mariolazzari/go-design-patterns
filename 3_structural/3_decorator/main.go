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
