package main

import (
	"fmt"
	"log"

	"aniket-batabyal.com/greetings"
)

/*

	Learnings:

		1. Create maps, slices (Data structures)
		2. run loops over slices, for loops
		3. Handle errors in loops

*/

func main() {

	names := []string{
		"Aniket",
		"Animesh",
		"Meeta",
	}
	// message, error := greetings.GreetUser("Aniket")
	messages, error := greetings.GreetUsers(names)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(messages)
}
