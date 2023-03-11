package main

import "fmt"

const (
	// this creates 1 followed by 100 zeros
	BIG_INT = 1 << 100

	// this goes back 99 places to give 1 or 2
	SMALL_INT = BIG_INT >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(SMALL_INT))
	fmt.Println(needFloat(BIG_INT))
	fmt.Println(needFloat(BIG_INT))
}
