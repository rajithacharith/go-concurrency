package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 3) // Buffered channel with capacity 3
	var wg sync.WaitGroup

	// Multiple goroutines sending data
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: sending value %d\n", val, val)
			ch <- val // The 4th send will block until space is available
			fmt.Printf("Goroutine %d: sent value %d\n", val, val)
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Receiving from channel:")

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Printf("Main: received value %d\n", v)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("All goroutines finished")
}
