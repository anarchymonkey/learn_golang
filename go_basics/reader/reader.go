package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/tour/reader"
)

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

type MyReader struct{}

// Reader is an interface that contains Read method, it takes a byte stream and then recieves the reader by itself and returns the length of bytes and the error if any
func (m MyReader) Read(b []byte) (int, error) {

	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

func main() {
	r := strings.NewReader("Hey my name is Aniket!")
	b := make([]byte, 8)
	for {
		lengthOfBytes, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", lengthOfBytes, err, b)
		fmt.Printf("b[:n] = %q\n", b[:lengthOfBytes])

		if err == io.EOF {
			break
		}

	}

	// Custom reader implementation

	fmt.Println("Validating the reader functionality")
	reader.Validate(MyReader{})
}
