package mathematicalOperations

import (
	"testing"
)

func TestSummationNumbers(t *testing.T) {

	var valuesToSum []rune = []rune{
		1,
		2,
		3,
		4,
		5,
	}

	want := 15
	summation := Summation(valuesToSum)

	if rune(want) != summation {
		t.Fatalf("The summation values was not the expected result that we wanted")
	}
}

func TestMultiplicationNonZeroNumbers(t *testing.T) {
	var numbersToMultiply []int = []int{
		1,
		2,
		3,
		4,
		5,
	}

	want := 120

	multiplicationValue, error := Multiplication(numbersToMultiply)

	if error != nil && want != multiplicationValue {
		t.Fatal(error)
	}

}
