package main

import (
	"fmt"
	"math"
)

// An interface is a type which can hold any types that can implement the methods inside it

// let us go through an example

// Animal can be an interface, an animals

type Shape interface {
	calcArea() float64
}

type Square struct {
	length int
}

type Circle struct {
	radius int
}

type Rectangle struct {
	Width  int
	Height int
}

func printArea(s Shape) {
	fmt.Println("The area is", s.calcArea())
}

func main() {
	var s []Shape
	var a Shape

	var circle Circle = Circle{
		radius: 25,
	}

	var square Square = Square{
		length: 10,
	}

	// var rect Rectangle = Rectangle{
	// 	Width:  10,
	// 	Height: 100,
	// }

	var rect_v1 Rectangle = Rectangle{
		Width:  100,
		Height: 0,
	}

	// a = rect
	a = &rect_v1

	printArea(a)

	// adding multiple shapes so that we can actually encapsulate the logic
	s = []Shape{
		&square,
		&circle,
	}

	for _, shape := range s {
		printArea(shape)
	}

}

func (s *Square) calcArea() float64 {
	return float64(s.length) * float64(s.length)
}

func (c *Circle) calcArea() float64 {
	return float64(c.radius) * float64(c.radius) * math.Pi
}

func (r *Rectangle) calcArea() float64 {
	return float64(r.Width) * float64(r.Height)
}
