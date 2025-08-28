package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	doOnce := func() {
		fmt.Println("This should print only once!")
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d calling Once\n", id)
			once.Do(doOnce)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines finished")
}
