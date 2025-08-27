
# Starvation Example in Go

This example demonstrates a starvation scenario using two goroutines (workers) and a shared mutex in Go.

## What is Starvation?

Starvation occurs when one or more goroutines are perpetually denied access to resources because other goroutines are continuously acquiring those resources. The starved goroutine rarely gets a chance to run, even though it is ready and waiting.

## Code Overview

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    var sharedLock sync.Mutex
    const runtime = 1 * time.Second

    greedyWorker := func() {
        defer wg.Done()
        var count int
        for begin := time.Now(); time.Since(begin) <= runtime; {
            sharedLock.Lock()
            time.Sleep(3 * time.Nanosecond)
            sharedLock.Unlock()
            count++
        }
        fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
    }

    politeWorker := func() {
        defer wg.Done()
        var count int
        for begin := time.Now(); time.Since(begin) <= runtime; {
            sharedLock.Lock()
            time.Sleep(1 * time.Nanosecond)
            sharedLock.Unlock()
            sharedLock.Lock()
            time.Sleep(1 * time.Nanosecond)
            sharedLock.Unlock()
            sharedLock.Lock()
            time.Sleep(1 * time.Nanosecond)
            sharedLock.Unlock()
            count++
        }
        fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
    }

    wg.Add(2)
    go greedyWorker()
    go politeWorker()
    wg.Wait()
}
```

## How Starvation Happens Here

- The greedy worker acquires the lock, does a little work, and releases it, repeating as fast as possible.
- The polite worker tries to be fair by releasing the lock more often, but as a result, it gets less work done because the greedy worker keeps grabbing the lock.
- The output will show that the greedy worker completes far more work loops than the polite worker, demonstrating starvation.

## How to Run

From the `5. starvation` directory, run:
```sh
go run main.go
```
You will see output like:
```
Greedy worker was able to execute 100000+ work loops
Polite worker was able to execute 1000+ work loops.
```
The greedy worker's output will appear much more frequently than the polite worker's output.

## Learn More
- [Starvation (Wikipedia)](https://en.wikipedia.org/wiki/Starvation_(computer_science))
- [Go sync package](https://pkg.go.dev/sync)
