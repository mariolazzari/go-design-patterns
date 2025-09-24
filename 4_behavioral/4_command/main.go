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
