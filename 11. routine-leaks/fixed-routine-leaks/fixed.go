package main

import (
	"fmt"
)

func safeRoutine(ch chan int, done chan struct{}) {
	for v := range ch {
		fmt.Println("Received:", v)
	}
	fmt.Println("Safe routine exiting")
	done <- struct{}{}
}

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	go safeRoutine(ch, done)

	ch <- 1
	ch <- 2
	fmt.Println("Sent two values")

	close(ch) // Properly close the channel to signal goroutine to exit
	<-done    // Wait for goroutine to finish
	fmt.Println("Main exiting, no goroutine leak")
}
