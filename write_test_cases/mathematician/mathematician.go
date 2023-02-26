package main

import (
	"fmt"
	"log"

	"aniket-batabyal.com/mathematicalOperations"
)

func main() {

	var numbersToSum []rune = []rune{
		1,
		2,
		3,
		4,
		5,
	}

	var numbersToMultiply []int = []int{
		10,
		20,
		30,
		40,
		1,
		10,
	}

	multipliedValues, error := mathematicalOperations.Multiplication(numbersToMultiply)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("The summation result: ", mathematicalOperations.Summation(numbersToSum))
	fmt.Println("The multiplication result: ", multipliedValues)
}
