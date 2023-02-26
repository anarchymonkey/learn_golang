package mathematicalOperations

import (
	"errors"
)

func Summation(numbers []rune) rune {
	var sum rune = 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}

func Multiplication(numbers []int) (int, error) {
	var multipliedSoFar int = 1

	for _, num := range numbers {

		if num < 1 {
			return 0, errors.New("numbers to be multiplied should be greater than 0")
		}

		multipliedSoFar *= num
	}

	return multipliedSoFar, nil
}
