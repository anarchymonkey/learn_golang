package main

import (
	"fmt"
	"math"
)

/*

Define a function by using these standard syntaxes

func <name of your func>(arguments) <return type> {

}
*/

// this is the default way of defining functions
func add(a int, b int) int {
	return a + b
}

// arguments can be defined using a shorthand like this
func multiply(a, b int) int {
	return a * b
}

// This type of function can be exported out of a package scope
func DoMultipleOperations(ops string, argumentList []int) int {

	if ops == "*" {
		multiplier := 1
		for _, argument := range argumentList {
			multiplier = multiply(multiplier, argument)
		}
		return multiplier
	}

	if ops == "+" {
		sum := 0

		for _, argument := range argumentList {
			sum = add(sum, argument)
		}
		return sum
	}

	return -1
}

// closures contain a function inside a function so that it would act as a Higher Order Function
func showClosures(a int) func(int) (int, error) {
	return func(b int) (int, error) {
		return a + b, nil
	}
}

// functions can be passed as values too
func addWithPi(fn func(int) int, valueToAdd int) int {
	return fn(valueToAdd) + int(math.Round(math.Pi))
}

func functionsAsValues(nextElementToAdd int) int {
	const RANDOM_VALUE_BASE = 10
	funcAsValue := func(a int) int {
		return RANDOM_VALUE_BASE + a
	}

	return addWithPi(funcAsValue, nextElementToAdd)
}

func fibonacci(previousValue int) func() int {
	a, b := 0, 1
	fmt.Println("a and b at init state", a, b)
	return func() int {
		result := a
		a, b = b, a+b

		fmt.Println("a and b", a, b)

		return result
	}
}

func printFibonacci() {
	previousValue := 0
	fn := fibonacci(previousValue)

	for i := 1; i < 10; i++ {
		fmt.Println(fn())
	}

}

func main() {
	fmt.Println("The result of sumation of two numbers is")
	fmt.Println(add(1, 2))

	fmt.Println("The result of the multiplication of two numbers using function shorthand is")
	fmt.Println(multiply(10, 20))

	fmt.Println("The result of Named export function is")
	fmt.Println(DoMultipleOperations("*", []int{1, 2, 3, 4}))

	fmt.Println("The result of the closure function")
	fmt.Println(showClosures(1)(3))

	fmt.Println("Function as value")
	fmt.Println(functionsAsValues(24))

	fmt.Println("Fibonacci closure")
	printFibonacci()

}
