package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var counter int
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
