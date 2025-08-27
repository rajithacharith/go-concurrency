# sync.Cond Example in Go

This directory demonstrates the use of `sync.Cond` for goroutine coordination in Go.

## What is sync.Cond?

`sync.Cond` is a condition variable that allows goroutines to wait for or signal events. It is useful for coordinating state changes between goroutines, such as waiting for a resource to become available or for a condition to be met.

## Example

**File:** `main.go`

This example shows a waiter goroutine waiting for a condition to be signaled by the main goroutine.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var mu sync.Mutex
    cond := sync.NewCond(&mu)
    ready := false

    // Waiter goroutine
    go func() {
        mu.Lock()
        for !ready {
            fmt.Println("Waiter: waiting...")
            cond.Wait()
        }
        fmt.Println("Waiter: done waiting!")
        mu.Unlock()
    }()

    // Simulate some work before signaling
    time.Sleep(100 * time.Millisecond)
    mu.Lock()
    ready = true
    fmt.Println("Main: signaling condition")
    cond.Signal()
    mu.Unlock()

    // Give time for waiter to finish
    time.Sleep(100 * time.Millisecond)
}
```

**Output:**
```
Waiter: waiting...
Main: signaling condition
Waiter: done waiting!
```

## How to Run

From the `8. cond` directory, run:
```sh
go run main.go
```

## Summary
- Use `sync.Cond` to coordinate goroutines based on conditions.
- `Wait()` blocks until `Signal()` or `Broadcast()` is called.

## Learn More
- [sync.Cond documentation](https://pkg.go.dev/sync#Cond)
- [Go Concurrency Patterns](https://golang.org/doc/effective_go#concurrency)
