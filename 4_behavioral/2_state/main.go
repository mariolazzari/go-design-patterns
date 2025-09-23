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
