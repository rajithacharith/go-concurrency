package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10000; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}
	// Sleep to give goroutines time to finish (not reliable)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Main function finished")
}
