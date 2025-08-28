# Channels Example in Go

This directory demonstrates the use of channels for communication between goroutines in Go.

## What are Channels?

Channels are Go's concurrency primitives that allow goroutines to communicate and synchronize by sending and receiving values. Channels provide a safe way to share data between goroutines without explicit locks.

- **Unbuffered channels** block the sender until the receiver is ready, and vice versa.
- **Buffered channels** allow a limited number of values to be queued before blocking.

## Example

**File:** `main.go`

```go
package main

import (
    "fmt"
)

func main() {
    ch := make(chan string)

    go func() {
        ch <- "Hello from goroutine!"
    }()

    msg := <-ch
    fmt.Println(msg)
}
```

**Output:**
```
Hello from goroutine!
```

## How Channels Work
- The main goroutine creates a channel and starts a new goroutine.
- The new goroutine sends a message into the channel.
- The main goroutine receives the message from the channel and prints it.
- Channels synchronize the two goroutines, ensuring safe communication.

## How to Run

From the `10. channels` directory, run:
```sh
go run main.go
```

## Summary
- Use channels to safely communicate between goroutines.
- Channels synchronize execution and transfer data without explicit locks.

## Learn More
- [Go Channels documentation](https://golang.org/doc/effective_go#channels)
- [Go Concurrency Patterns](https://golang.org/doc/effective_go#concurrency)
