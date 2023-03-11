package main

import (
	"fmt"
	"runtime"
	"time"
)

func getExampleOfABasicIfStatement(fathersAge int, otherPersonsAge int) string {

	if fathersAge > otherPersonsAge {
		return "Luke Skywalker, I am your father!"
	}
	return "Welcome to your doom, I am Darth Vader"
}

func getExampleOfIfWithAShortSyntax(dailySpentAmount int, limit int) string {

	if v := dailySpentAmount * 30; v > limit {
		return "You have crossed your daily limit"
	} else {
		fmt.Println("Good job saving some money")
	}

	return "Keep on spending"
}

func getExamplesOfSwtichCases() string {
	// lets learn about switch cases

	fmt.Println("Go runs on")

	osRuntime := runtime.GOOS
	// does not need break cause it automatically puts it
	switch osRuntime {
	case "darwin":
		return "This is the darwin Runtime"
	case "linux":
		return "Linux"
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", runtime.GOOS)
		return ""
	}
}

func getSpecialExamplesOfSwitch() {
	fmt.Println("\nWhen is saturday")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("The next day")
	case today + 2:
		fmt.Println("T + 2 days")
	default:
		fmt.Println("It's too far, cant calculate")
	}
}

func main() {
	// basic if statement
	fmt.Println(getExampleOfABasicIfStatement(50, 25))

	// short if statement

	fmt.Println(getExampleOfIfWithAShortSyntax(100, 20000))

	// switch statements

	getExamplesOfSwtichCases()

	// more examples

	getSpecialExamplesOfSwitch()

}
