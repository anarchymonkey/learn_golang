package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

type Address struct {
	city            string
	state           string
	basicAddress    string
	extendedAddress string
	pincode         int
}
type UserProfile struct {
	name     string
	nickName string

	age int

	address Address
}

func WordCount(s string) (map[string]int, error) {

	if !utf8.ValidString(s) || s == "" {
		return nil, errors.New("Cannot parse non utf-8 embedded string")
	}
	wordCount := make(map[string]int)

	for _, value := range strings.Split(s, " ") {
		if _, hasKey := wordCount[string(value)]; hasKey {
			wordCount[string(value)] = wordCount[string(value)] + 1
		} else {
			wordCount[string(value)] = 1
		}
	}

	return wordCount, nil
}

var m map[string]UserProfile

func GetUserProfile() {
	// What are maps?
	// Ans: Maps are just a store of keys and values

	m = make(map[string]UserProfile)

	m["Aniket"] = UserProfile{
		name:     "Aniket",
		nickName: "Rick",
		age:      27,
		address: Address{
			city:            "Bangalore",
			state:           "Karnataka",
			basicAddress:    "Jhingalala huhu",
			extendedAddress: "huhu huhu",
			pincode:         560001,
		},
	}

	fmt.Println("Maps")
	fmt.Println("Map: ", m["Aniket"])

	// map literals

	var mapLiteral map[string]UserProfile = map[string]UserProfile{
		"Meeta": {
			name:     "Meeta",
			nickName: "Meetu",
			age:      27,
			address: Address{
				city:            "Bangalore",
				state:           "Karnataka",
				basicAddress:    "Jhingalala huhu",
				extendedAddress: "huhu huhu",
				pincode:         560075,
			},
		},
		"Aniket": {
			name:     "Aniket",
			nickName: "Rick",
			age:      26,
			address: Address{
				city:            "Bangalore",
				state:           "Karnataka",
				basicAddress:    "Jhingalala huhu",
				extendedAddress: "huhu huhu",
				pincode:         560015,
			},
		},
	}

	for key, userProfile := range mapLiteral {
		fmt.Printf("Nickname at index = %s is %v \n", key, userProfile.address)
	}

	aniketsDetails := mapLiteral["Aniket"]

	fmt.Println("Anikets details are ", aniketsDetails)

	// delete a key

	delete(mapLiteral, "Meeta")

	// check if key exists or not

	elem, ok := mapLiteral["Aniket"]

	if ok {
		fmt.Println("The values are", elem, ok)
	}
}

func main() {
	GetUserProfile()

	fmt.Println("Word count map is ")
	fmt.Println(WordCount("Hey my name is Aniket"))
	fmt.Println(WordCount("Aniket is awesome, Aniket is so good"))
}
