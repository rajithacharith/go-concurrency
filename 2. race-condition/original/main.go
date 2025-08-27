package main

import (
	"fmt"
	"time"
)

func main() {
	var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++ // Race condition: multiple goroutines access counter
		}()
	}
	time.Sleep(1 * time.Second) // Wait for goroutines to finish
	fmt.Println("Final counter value:", counter)
}
