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

type DummyInterface interface {
	validateSquare()
}

func typeAssertions() {
	var inter interface{} = "hello"

	defer func() {
		// this recoveres the panic function if any
		println("Recovered", recover())
	}()

	s, ok := inter.(string)

	if ok {
		fmt.Println("interface type assertion", s)
	}

	f, floatOk := inter.(float64)

	if floatOk {
		fmt.Println("Suddenly the type changed to ", f)
	} else {
		fmt.Println("Cannot implement other types than the defined one")
	}

}

func typeSwitches(i interface{}) {

	switch v := i.(type) {
	case *Square:
		{
			fmt.Println("This is a square", v.length)
		}
	case *Circle:
		{
			fmt.Println("This is a circle", v.radius, v.calcArea())
		}
	default:
		{
			fmt.Println("This is some unknown type", v)
		}
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
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

	// interface values with nil underlying values
	var nullSquare *Square = &Square{
		length: 20,
	}
	var b DummyInterface = nullSquare

	b.validateSquare()

	// type assertion

	typeAssertions()

	// typeswitches

	for _, shape := range s {
		typeSwitches(shape)
	}

	// stringer

	fmt.Println(&square)

	// stringer exercise

	var ipAddrStore map[string]IPAddr = map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for _, value := range ipAddrStore {
		fmt.Println(value)
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

func (s *Square) validateSquare() {

	if s == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println("Its a validated square")
}

func (s *Square) String() string {
	return fmt.Sprintf("The square value is %d", s.length)
}
