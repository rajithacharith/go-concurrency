package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	fmt.Println("Starting goroutines...")

	for i := 1; i <= 5; i++ {
		go worker(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("All workers completed")
}
