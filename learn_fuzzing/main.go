package main

import (
	"errors"
	"fmt"
	"log"
	"unicode/utf8"
)

func ReverseString(str string) (string, error) {

	if !utf8.ValidString(str) {
		return "", errors.New("The string is not valid UTF-8 %s" + str)
	}
	b := []rune(str)

	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b), nil
}

func main() {

	const NEW_STRING = "Aniket"

	reversedString, error := ReverseString(NEW_STRING)

	if error != nil {
		log.Fatal("Some error occured")
	}

	doubleReversedString, error := ReverseString(reversedString)

	if error != nil {
		log.Fatal("Some error occured")
	}

	fmt.Println("The reversed string is", reversedString)
	fmt.Println("The double reversed string is", doubleReversedString)
}
