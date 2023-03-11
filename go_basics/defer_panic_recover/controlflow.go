package main

import "fmt"

// We will learn about defer, panic, recover

func calcSum[K int | float64](a K, b K) K {
	return a + b
}

func deferFuncExample() {
	// defer stalls the execution of the instance linked to it
	// defer runs after the corrosponding function block has finished running
	// multiple defers use the LIFO algo, last defer goes out first
	defer fmt.Println("Something has been calculated")
	sum := calcSum(1, 2)
	fmt.Println("The sum is", sum)

	defer fmt.Println("The sum of floats has been calculated")

	fmt.Println(calcSum(1.100, 2.200))
}

func panicFuncExample(randomValue int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The function has recovered", r)
		}
	}()

	fmt.Println("Before the panic happened")
	if randomValue == 10 {
		panic(fmt.Sprintf("returned value is %d", randomValue))
	}

	fmt.Println("Normal execution has started")
}

func main() {
	deferFuncExample()

	// panic and recover example

	panicFuncExample(10)
}
