package main

import (
	"fmt"
)

// GO has only for loop

/*
	Parts of for loop

	the init statement: executed before the first iteration

		this statement would be a short variable declaration and it is visible only to the scope of the for statement

	the condition of iteration: evaluated before every iteration

		The loop will stop iteration when the condition becomes false

	the post statement: executed at the end of every iteration

*/

func getSimpleForLoopExample() {
	var sum int

	// i := 0 inits the value of i to 0
	// i < 10 says that keep the loop going until the value comes to 10 and condition is true
	// i++ is the post statement, it will increment i everytime one iteration ends

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("The sum is", sum)
}

func getNonInitAndPostForLoop() {

	// Another modification is that the init and post statements are not always required
	// lets just update our whole programme with a single summation loop
	var sum int = 1

	// can be understood as a while loop
	for sum < 1000 {
		sum += sum
	}

	fmt.Println("The non init and post sum is", sum)

}

// range is a short form of an iterator used to iterate maps, sets, slices
func getRangeFormOfLoops() {

	primeNumbers := []int{
		1,
		2,
		3,
		5,
		7,
		11,
		13,
	}

	// we can skip index, value by using _ as the underlying value
	for idx, value := range primeNumbers {
		fmt.Printf("{index, value} = {%v, %v} ", idx, value)
	}
}

func main() {

	// simple for loop

	getSimpleForLoopExample()

	// non init and post for loop

	getNonInitAndPostForLoop()

	// infinite loop example

	/*
		for {

		}
	*/

	// get Range form loops
	getRangeFormOfLoops()
}
