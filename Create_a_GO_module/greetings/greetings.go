package greetings

import "fmt"

// function names starting with Capital letters are called exported names

func GreetUser(name string) string {
	message := fmt.Sprintf("Hey, welcome to my channel %v", name)
	return message
}
