package main

import "testing"

func TestWordCount(t *testing.T) {

	actual_output, error := WordCount("Hey my name is Aniket")

	if error != nil {
		t.Error("Found error", error)
		return
	}

	newMap := make(map[string]int)

	newMap = map[string]int{
		"Aniket": 1,
		"is":     1,
		"name":   1,
		"my":     1,
		"Hey":    1,
	}

	for key, value := range actual_output {
		keyCountInNewMap := newMap[key]
		actualOutputCount := value

		if actualOutputCount != keyCountInNewMap {
			t.Error("The map count does not match")
		}
	}
}

func TestWordCountEmpty(t *testing.T) {

	_, error := WordCount("")

	if error == nil {
		t.Error("Found error", error)
		return
	}
}
