package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false

	// Waiter goroutine
	go func() {
		mu.Lock()
		for !ready {
			fmt.Println("Waiter: waiting...")
			cond.Wait()
		}
		fmt.Println("Waiter: done waiting!")
		mu.Unlock()
	}()

	// Simulate some work before signaling
	time.Sleep(100 * time.Millisecond)
	mu.Lock()
	ready = true
	fmt.Println("Main: signaling condition")
	cond.Signal()
	mu.Unlock()

	// Give time for waiter to finish
	time.Sleep(100 * time.Millisecond)
}
