package main

import "fmt"

func getDefaultValuesOfDataTypesExamples() {
	var integerValue int
	var stringValue string
	var booleanValue bool

	fmt.Printf("Integer value = %v, String value = %v, Boolean value = %v \n", integerValue, stringValue, booleanValue)
}

func getTypeConversionExamples() {
	var integerValue int = -12
	var floatValue float32 = float32(integerValue) * 0.1234 * -1
	var unsignedIntValue uint32 = uint32(floatValue)

	fmt.Printf(
		"The integer value = %d, float value = %v, Unsigned int value = %v \n",
		integerValue, floatValue, unsignedIntValue,
	)
}

func getTypeInferrenceExamples() {
	var i int = 10

	j := i

	fmt.Printf("The type of J is %T", j)
}

func main() {

	fmt.Println("Hey")

	// default types

	getDefaultValuesOfDataTypesExamples()

	// type conversion

	getTypeConversionExamples()

	// type infferrence

	getTypeInferrenceExamples()
}
