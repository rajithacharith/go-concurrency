package main

import (
	"fmt"
	"sync"
	"time"
)

type stick struct{ sync.Mutex }

func main() {
	var left, right stick
	done := make(chan struct{})

	eat := func(name string, first, second *stick) {
		for i := 0; i < 10; i++ {
			first.Lock()
			time.Sleep(1 * time.Millisecond) // Try to avoid deadlock
			if !second.TryLock() {
				first.Unlock()
				fmt.Printf("%s: couldn't acquire both sticks, retrying\n", name)
				time.Sleep(1 * time.Millisecond)
				continue // Livelock: both keep retrying
			}
			fmt.Printf("%s: eating\n", name)
			second.Unlock()
			first.Unlock()
			time.Sleep(1 * time.Millisecond)
		}
		done <- struct{}{}
	}

	go eat("A", &left, &right)
	go eat("B", &right, &left)
	<-done
	<-done
}
