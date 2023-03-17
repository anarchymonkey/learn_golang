package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

/*
	Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/
type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	length, err := rot.r.Read(b)

	if err != nil {
		return 0, errors.New("Some error occured while reading the file")
	}

	for i := 0; i < length; i++ {
		if (b[i] >= 'A' && b[i] <= 'M') || (b[i] >= 'a' && b[i] <= 'm') {
			b[i] = b[i] + 13
		} else if (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'n' && b[i] <= 'z') {
			b[i] = b[i] - 13
		}
	}

	return length, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := &rot13Reader{s}
	io.Copy(os.Stdout, r)
}
