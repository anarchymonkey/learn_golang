package main

import (
	"fmt"

	"batsy.com/utils"
)

func main() {
	fmt.Println("This is a simple example of a generics module")

	integerValues := map[string]int{
		"first":  100,
		"second": 200,
	}

	floatingValues := map[string]float64{
		"first":  100.20,
		"second": 200.80,
	}

	// fmt.Println("as we can see that we need to make a lot of functions to get the same summation of numbers")
	// fmt.Println(utils.SumOfNIntegers(integerValues))
	// fmt.Println(utils.SumOfNFloats(floatingValues))

	fmt.Println("Now by using generics we can use the same function for the calculation of every number")

	fmt.Println(utils.SumOfNNumbers(integerValues))
	fmt.Println(utils.SumOfNNumbers(floatingValues))
}
