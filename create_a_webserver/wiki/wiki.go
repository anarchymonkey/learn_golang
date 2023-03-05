package wiki

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

const EXTENTION = "txt"

// This is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error.
func (p *Page) Save() error {
	fileName := fmt.Sprintf("%s.%s", p.Title, EXTENTION)

	// 0600 means read write permissions
	return os.WriteFile(fileName, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	fileName := fmt.Sprintf("%s.%s", title, EXTENTION)
	body, error := os.ReadFile(fileName)

	if error != nil {
		return nil, error
	}
	return &Page{Title: title, Body: body}, nil
}
