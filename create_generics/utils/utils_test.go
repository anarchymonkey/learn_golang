package utils

import (
	"log"
	"testing"
)

func TestSumOfNIntegers(t *testing.T) {
	integerNumbers := map[string]int{
		"first":  10,
		"second": 20,
	}
	actual_result := SumOfNIntegers(integerNumbers)
	expected_result := 30

	if actual_result != expected_result {
		log.Fatalf("Some error occured, wrong result")
	}
}

func TestSumOfNFloats(t *testing.T) {

	var floatingNumbers map[string]float64 = map[string]float64{
		"first":  10.01,
		"second": 20.09,
	}

	actual_result := SumOfNFloats(floatingNumbers)

	expected_result := 30.10

	if actual_result != expected_result {
		log.Fatalf("Floating number results do not match")
	}

}

func TestSumOfNNumbers(t *testing.T) {
	integerNumbers := map[string]int{
		"first":  10,
		"second": 20,
	}

	actual_result := SumOfNNumbers(integerNumbers)
	expected_result := 30

	if actual_result != expected_result {
		log.Fatal("The results do not match")
	}
}
