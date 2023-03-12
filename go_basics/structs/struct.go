package main

import "fmt"

// A struct is a collection of fields

type Vertex struct {
	X int
	Y int
}

func structExample() {

	// basic struct example
	var vertex1 Vertex = Vertex{
		X: 1,
		Y: 2,
	}

	fmt.Printf("X = %d, Y = %d \n", vertex1.X, vertex1.Y)

	// pointer to a struct example

	var vertex2 Vertex = Vertex{
		X: 10,
		Y: 20,
	}

	p := &vertex2

	// you can ommit the * operator while assigning values to the pointer
	p.X = 12

	fmt.Println("X in pointer is ", p.X)

	// struct literals

	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}
	v3 := Vertex{}
	ptr := &Vertex{3, 4}

	fmt.Println(v1, v2, v3, ptr.X, ptr.Y)

}

func main() {

	structExample()
}
