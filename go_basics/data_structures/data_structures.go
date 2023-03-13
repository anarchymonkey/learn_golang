package main

import (
	"fmt"
	"strings"
)

func learnArrays() {
	// arrays can be defined as the syntax [n]T
	// n denotes the size of the array whereas T defines the type

	var numbers [5]int

	for i := 0; i < 5; i++ {
		numbers[i] = i * 2
	}

	var primeNumbers [5]int = [5]int{1, 3, 5, 7, 11}

	fmt.Println(numbers, primeNumbers)

}

func learnDynamicArrays() {
	// a basic array turned into a slice

	randomMonthsInFeb := []int{
		1,
		2,
		3,
		4,
		19,
		20,
		28,
	}

	var sliceOfMonthsInFeb []int = randomMonthsInFeb[0:5]

	fmt.Println("The slice is", sliceOfMonthsInFeb)

	// appending values in slices

	for i := 0; i < 5; i++ {
		// remember append returns the value and thus needs to be stored somewhere
		sliceOfMonthsInFeb = append(sliceOfMonthsInFeb, i*4)
	}

	fmt.Println("appended slice", sliceOfMonthsInFeb)

	// A slice does not store any data, rather it just references an array location thus if we modify that location the copies will also be modified

	// Let us see with an example

	cities := []string{
		"Bangalore",
		"Hydrabad",
		"Kolkata",
		"Mumbai",
	}

	fmt.Println("cities => start", cities)

	a := cities[0:2]
	b := cities[1:3]

	fmt.Println("a & b", a, b)

	b[0] = "+++++++"

	// you will see that if we modify the reference of b[0], the same modification will be done in a[1] as they share the same reference

	fmt.Println("a & b after b is moditfied", a, b)

	fmt.Println("cities => end", cities)

	// slice defaults

	// lower bound is defaulted to 0 & the higher bound is defaulted to the length slice[lowerbound:upperbound]

	// slice length and capacity

	// empty slice is init with nil

	dummySlice := []int{
		1, 2, 3, 4, 5,
	}

	printSlice(dummySlice)

	dummySlice = dummySlice[0:3]

	printSlice(dummySlice)

	dummySlice = dummySlice[:2]

	printSlice(dummySlice)

	// Creating a slice with make

	fmt.Println("Creating slices with make, lets see the length and capacity")

	// assigns a length of 5
	newSlice := make([]int, 5)

	printSlice(newSlice)

	// assigns a capacity of 5
	newSlice = make([]int, 0, 5)

	newSlice = newSlice[:cap(newSlice)]
	newSlice = newSlice[0:3]

	printSlice(newSlice)

}

func createSlicesOfSlices() {
	// slices of slices in DSA terms is a 2D dynamic array

	// lets make a chess board

	chessBoard := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	chessBoard[0][0] = "X"
	chessBoard[1][1] = "O"
	chessBoard[0][1] = "X"
	chessBoard[2][2] = "Aniket"
	chessBoard[2][1] = "O"
	chessBoard[0][2] = "O"
	chessBoard[2][0] = "X"
	chessBoard[1][0] = "O"
	chessBoard[1][2] = "X"

	for i := 0; i < len(chessBoard); i++ {
		fmt.Printf("%s\n", strings.Join(chessBoard[i], " "))
	}
}

func appendThingsToSlices() {
	fmt.Println("This is appending someting to slices")

	someArray := [5]string{
		"Aniket",
		"Animesh",
		"Ankush",
		"Ankit",
		"Amaan",
	}

	someSlice := make([]string, 0)

	for i := 0; i < len(someArray); i++ {
		someSlice = append(someSlice, someArray[i])
	}

	someSlice = append(someSlice, "Adding name: Amitava")
	fmt.Println("Slice after appending")
	fmt.Println(someSlice)
	printSlice(someSlice)
}

func printSlice[T []int | []string](val T) {
	fmt.Printf("\n len = %v, cap = %v \n", len(val), cap(val))
}

func main() {

	// learn arrays
	learnArrays()

	// learn dynamic arrays | slices

	learnDynamicArrays()

	// create 2D arrays

	createSlicesOfSlices()

	// appendThingsToSlices

	appendThingsToSlices()
}
