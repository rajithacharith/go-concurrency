package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rw sync.RWMutex
	var value int
	var wg sync.WaitGroup

	// Writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			rw.Lock()
			value = i
			fmt.Printf("Writer set value to %d\n", value)
			rw.Unlock()
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Reader goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				rw.RLock()
				fmt.Printf("Reader %d read value %d\n", id, value)
				rw.RUnlock()
				time.Sleep(5 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}
