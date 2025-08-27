package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++ // No mutex: race condition
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
