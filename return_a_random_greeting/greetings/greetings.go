package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func GreetUser(name string) (string, error) {

	if name == "" {
		return "", errors.New("the name cannot be empty or nil")
	}
	message := fmt.Sprintf(randomHello(), name)
	return message, nil
}

func GreetUsers(names []string) (map[string]string, error) {

	// make the messages map by using the make function
	messages := make(map[string]string)

	// loop through the names slice
	for idx, name := range names {
		fmt.Println("Index is: ", idx)

		// call greet user and get the message and the error
		message, err := GreetUser(name)

		// if error is nil then return the error
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomHello() string {

	formattedHellos := []string{
		"Hey mate! %v",
		"Hola %v",
		"konichiwa %v",
	}

	// fmt.Println("The length of the formatted hellos screen is", len(formattedHellos))

	return formattedHellos[rand.Intn(len(formattedHellos))]
}
