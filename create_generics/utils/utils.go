package utils

// This is the union type moved to a new type
type Number interface {
	int | int64 | float64
}

func SumOfNIntegers(numbers map[string]int) int {
	var sum int = 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumOfNFloats(numbers map[string]float64) float64 {

	var sum float64 = 0.00

	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumOfNNumbers[K comparable, V Number](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}
