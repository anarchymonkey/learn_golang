package main

import "fmt"

func main() {
	fmt.Println("Let's learn about variables")

	// Declaring a variable
	var name string = "Aniket"

	// shorthand of declearing a variable
	nameShorthand := "Aniket"

	// declaring constants
	const PREFIX = "My name is"

	// declaring multiple variables at a go with default values

	var uno, dos, tres = 1, 2, 3

	fmt.Printf("name = %s, shortHand = %s, PREFIX = %s \n", name, nameShorthand, PREFIX)

	fmt.Printf("Uno = %d, Dos = %d, Tres = %d", uno, dos, tres)

}
