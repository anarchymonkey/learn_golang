package main

import "fmt"

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
func main() {
	fmt.Println("The result of sumation of two numbers is")
	fmt.Println(add(1, 2))

	fmt.Println("The result of the multiplication of two numbers using function shorthand is")
	fmt.Println(multiply(10, 20))

	fmt.Println("The result of Named export function is")
	fmt.Println(DoMultipleOperations("*", []int{1, 2, 3, 4}))

}
