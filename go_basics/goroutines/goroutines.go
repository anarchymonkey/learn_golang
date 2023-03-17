package main

import (
	"fmt"
	"time"
)

func saySomething(text string) {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s \n", text)
	}

}

func main() {

	// side villans
	go saySomething("Meeta is cute")

	// main villan
	saySomething("When you run it what will happen?")
}
