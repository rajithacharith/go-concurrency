package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines finished")
}
