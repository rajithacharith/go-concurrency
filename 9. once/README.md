# sync.Once Example in Go

This directory demonstrates the use of `sync.Once` to ensure a function is executed only once, even when called from multiple goroutines.

## What is sync.Once?

`sync.Once` is a concurrency primitive that guarantees a function will only be executed once, no matter how many times its `Do` method is called, even across multiple goroutines. This is useful for one-time initialization (e.g., opening a file, initializing a resource).

## Example

**File:** `main.go`

```go
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
```

**Output:**
```
Goroutine 0 calling Once
Goroutine 1 calling Once
Goroutine 2 calling Once
Goroutine 3 calling Once
Goroutine 4 calling Once
This should print only once!
All goroutines finished
```

## How to Run

From the `9. once` directory, run:
```sh
go run main.go
```

## Summary
- Use `sync.Once` for safe one-time initialization in concurrent programs.
- Only the first call to `Do` will execute the function; all others are ignored.

## Learn More
- [sync.Once documentation](https://pkg.go.dev/sync#Once)
- [Go Concurrency Patterns](https://golang.org/doc/effective_go#concurrency)
