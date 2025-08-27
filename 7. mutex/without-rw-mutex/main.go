package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// No RWMutex used here
	var value int
	var wg sync.WaitGroup

	// Writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			value = i
			fmt.Printf("Writer set value to %d\n", value)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Reader goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				fmt.Printf("Reader %d read value %d\n", id, value)
				time.Sleep(5 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}
