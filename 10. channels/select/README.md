# Select Statement Example in Go

This directory demonstrates how to use the `select` statement to handle multiple channels in Go.

## What is the Select Statement?

The `select` statement lets a goroutine wait on multiple communication operations. The first channel that is ready will proceed, allowing you to handle whichever event happens first. This is useful for multiplexing input from several sources, implementing timeouts, and more.

## Example

**File:** `main.go`

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(500 * time.Millisecond)
        ch1 <- "Message from channel 1"
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "Message from channel 2"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        }
    }
}
```

**Output:**
```
Message from channel 1
Message from channel 2
```

## How It Works
- Two goroutines send messages to two different channels after different delays.
- The main goroutine uses a `select` statement to receive from whichever channel is ready first.
- The output order depends on which message arrives first.

## How to Run

From the `select` directory, run:
```sh
go run main.go
```

## Summary
- Use `select` to wait on multiple channels and handle whichever is ready first.
- Useful for multiplexing, timeouts, and concurrent event handling.

## Learn More
- [Go Select Statement documentation](https://golang.org/doc/effective_go#select)
- [Go Concurrency Patterns](https://golang.org/doc/effective_go#concurrency)
