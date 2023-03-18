package main

import (
	"fmt"
	"time"
)

func sum(nums []int, c chan int) {
	length := len(nums)
	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
	}

	fmt.Println("Sending to channel")
	c <- sum
}

func fibonacci(num int, c chan int) {
	x, y := 0, 1

	for i := 0; i < num; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func fibonacciWithSelectOps(c, quit chan int) {

	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quitting")
			return
		}
	}
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

	// range and close

	var fiboChannel chan int = make(chan int, 10)

	go fibonacci(cap(fiboChannel), fiboChannel)

	// this will recieve values untill the channel is closed
	for i := range fiboChannel {
		fmt.Println(i)
	}

	// ******************* END ************************* //

	// select operations

	// a select operation lets a goroutine wait on multiple communication operations

	selectChan, quit := make(chan int, 15), make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("val", <-selectChan)
		}
		quit <- 0
	}()

	fibonacciWithSelectOps(selectChan, quit)

	// select statement with default

	fmt.Println("\n")

	timer := time.Tick(100 * time.Millisecond)
	boomer := time.After(500 * time.Millisecond)

	// default runs if the channels are not ready (i.e the cases inside are not reade)

	// here the timer runs after every 100 * milisecond delta
	// boom runs after every 500 miliseconds * delta
	// before these two are ready the default case runs and sleeps the programme for 25 * delta miliseconds (1/4 th of timer)
	for {
		select {
		case <-timer:
			fmt.Println("Ticking")
		case <-boomer:
			fmt.Println("Boom! The programme exploded")
			return
		default:
			fmt.Println("....")
			time.Sleep(25 * time.Millisecond)
		}
	}

}
