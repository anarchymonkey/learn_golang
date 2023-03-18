package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Unique struct {
	mu          sync.Mutex
	valueMapper map[int]int
}

func (u *Unique) generateUniqueIdCollector(index int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("index is", index)
	u.mu.Lock()
	u.valueMapper[index] = int(rand.Intn(255))
	u.mu.Unlock()

}

func main() {

	var uniq Unique = Unique{
		valueMapper: make(map[int]int),
	}

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go uniq.generateUniqueIdCollector(i, &wg)
	}

	wg.Wait()
	// time.Sleep(time.Second)
	fmt.Println(uniq.valueMapper)
	// wg.Done()
}
