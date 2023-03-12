package main

import "fmt"

// note pointers cannot be assigned to addressess of other datatypes
func main() {
	fmt.Println("Pointers example")

	name, age := "Aniket", 27

	p := &name

	fmt.Println("pointer value p is = ", *p)

	*p = "Animesh"

	fmt.Println("Changed value of name is", name)

	dob := "27th April, 1996"

	p = &dob

	fmt.Println("pointer after reassignment is", *p)

	fmt.Printf("Finally my name is = %s, my dob is = %s, age is = %d", name, *p, age)
}
