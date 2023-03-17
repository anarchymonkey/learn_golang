package main

import "fmt"

func printValues[T comparable](values []T) {
	fmt.Println("The values are", values)
}

type DataParser[T any, I int] struct {
	Data []T

	Name string

	Id I
}

func main() {

	fmt.Println("learning about type Parameters")

	s := []string{
		"Aniket",
		"Amitava",
		"Tula",
	}

	nums := []int{
		1, 2, 3, 4, 5,
	}

	printValues(s)
	printValues(nums)

	// generics

	var parsedData DataParser[string, int] = DataParser[string, int]{
		Data: s,
		Name: "Aniket",
		Id:   1,
	}

	fmt.Printf("The data is %v & the name to which it belongs to is %s, the id is %d", parsedData.Data, parsedData.Name, parsedData.Id)

}
