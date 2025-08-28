package main

import (
	"fmt"
	"time"
)

func leakyRoutine(ch chan int) {
	for v := range ch {
		fmt.Println("Received:", v)
	}
	// If channel is never closed, this goroutine will never exit
	fmt.Println("Leaky routine exiting")
}

func main() {
	ch := make(chan int)
	go leakyRoutine(ch)

	ch <- 1
	ch <- 2
	fmt.Println("Sent two values")

	// No close(ch), so leakyRoutine will block forever
	time.Sleep(1 * time.Second)
	fmt.Println("Main exiting, but leakyRoutine is still running (leak)")
}
