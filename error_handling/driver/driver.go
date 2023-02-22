package main

import (
	"fmt"

	"aniket-batabyal.com/car"
)

func main() {
	fmt.Println("Hey I am a driver")
	message, error := car.StartCar("Ford Mustang")

	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("The message is: ", message)
}
