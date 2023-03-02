package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

// This is a normal test case
func TestReverseString(t *testing.T) {

	testCases := []struct {
		in, want string
	}{
		{"Aniket", "tekinA"},
		{"hie", "eih"},
		{"Hue", "euH"},
	}

	for _, tc := range testCases {
		reversedString, error := ReverseString(tc.in)

		if error != nil {
			t.Errorf("Something Went Wrong")
		}

		if reversedString != tc.want {
			t.Errorf("The input was %s, the reversed string which did not match was %s, which was expected to be %s", tc.in, reversedString, tc.want)
		}
	}
}

func FuzzReverseString(f *testing.F) {

	// add a fuzz test case

	testcases := []string{
		"Aniket",
		"Namaste",
		"Jhingalala",
		// "",
	}

	for _, tc := range testcases {
		// This adds a seed corpus, in simple terms I think it takes a sample input which it can make permutations of
		f.Add(tc)
	}

	// This represents the actual testing scenario where we would run the test
	f.Fuzz(func(t *testing.T, originalString string) {
		reversedString, error := ReverseString(originalString)

		if error != nil {
			fmt.Println(error)
			return
		}
		doubleReverse, error := ReverseString(reversedString)

		if error != nil {
			t.Errorf("Some error occured in double Reverse string")
			return
		}

		if originalString != doubleReverse {
			t.Errorf("The original string reference did not match")
		}

		if utf8.ValidString(originalString) && !utf8.ValidString(reversedString) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", reversedString)
		}
	})
}
