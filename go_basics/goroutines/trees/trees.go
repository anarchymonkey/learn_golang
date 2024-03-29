package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {

	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {

	ch1, ch2 := make(chan int, 20), make(chan int, 20)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {

		v1, v2 := <-ch1, <-ch2

		if v1 != v2 {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)

	t2 := tree.New(1)

	fmt.Println(Same(t1, t2))

}
