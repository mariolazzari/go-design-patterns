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
