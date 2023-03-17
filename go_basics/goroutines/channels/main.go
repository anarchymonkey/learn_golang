package main

import "fmt"

func sum(nums []int, c chan int) {
	length := len(nums)
	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
	}

	fmt.Println("Sending to channel")
	c <- sum
}

func main() {

	nums := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	c := make(chan int)

	// this helps goroutines to sync
	go sum(nums[:len(nums)/2], c)
	go sum(nums[len(nums)/2:], c)

	firstSum, secondSum := <-c, <-c

	fmt.Printf("The first sum  = %d \n The second sum = %d \n", firstSum, secondSum)

	// buffered channels

	var channel_v1 chan int = make(chan int, 2)

	channel_v1 <- 1
	channel_v1 <- 2

	fmt.Println("first", <-channel_v1)
	fmt.Println("second", <-channel_v1)
}
