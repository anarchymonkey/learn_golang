package main

import (
	"fmt"

	"log"

	"aniket-batabyal.com/car"
)

func main() {
	// This says that set a prefix everytime you land in an error
	log.SetPrefix("######====> ")
	// This disables other info while we get errors like timestamp and all
	log.SetFlags(0)

	message, error := car.StartCar("")

	if error != nil {
		// this prints the error and exits the application by setting os.Exit(1)
		log.Print("A lot of errors is not good for your health")
		log.Fatal(error)
	}
	fmt.Println(message)
}
