# Buffered Channels Example in Go

This directory demonstrates the use of buffered channels for communication between multiple goroutines in Go.

## What are Buffered Channels?

Buffered channels allow you to send a limited number of values without blocking, up to the buffer's capacity. Once the buffer is full, further sends will block until space is available (i.e., until a value is received from the channel).

## Example

**File:** `main.go`

```go
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
```

**Output:**
```
Goroutine 1: sending value 1
Goroutine 2: sending value 2
Goroutine 3: sending value 3
Goroutine 4: sending value 4
Goroutine 1: sent value 1
Goroutine 2: sent value 2
Goroutine 3: sent value 3
Receiving from channel:
Main: received value 1
Main: received value 2
Main: received value 3
Goroutine 4: sent value 4
Main: received value 4
All goroutines finished
```

## How It Works
- The first three goroutines send values to the channel and succeed immediately (buffer not full).
- The fourth goroutine blocks until the main routine receives a value, freeing buffer space.
- The main routine receives all values using a `for v := range ch` loop, which continues until the channel is closed.
- The channel is closed after all goroutines finish sending.

## How to Run

From the `buffered-channels` directory, run:
```sh
go run main.go
```

## Summary
- Buffered channels allow multiple values to be sent without blocking, up to their capacity.
- When the buffer is full, further sends block until space is available.
- Use a `for v := range ch` loop to receive all values until the channel is closed.

## Learn More
- [Go Channels documentation](https://golang.org/doc/effective_go#channels)
- [Go Concurrency Patterns](https://golang.org/doc/effective_go#concurrency)
