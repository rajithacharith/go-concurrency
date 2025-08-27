# WaitGroup Example in Go

This directory demonstrates the difference between using and not using `sync.WaitGroup` to wait for goroutines to finish in Go.

## With WaitGroup

**File:** `with-wait-group/main.go`

This example uses a `sync.WaitGroup` to ensure the main function waits for all goroutines to finish before exiting.

```go
var wg sync.WaitGroup
for i := 1; i <= 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Printf("Goroutine %d finished\n", id)
    }(i)
}
wg.Wait()
fmt.Println("All goroutines finished")
```

**Output:**
```
Goroutine 1 finished
Goroutine 2 finished
...
All goroutines finished
```

## Without WaitGroup

**File:** `without-wait-group/main.go`

This example does not use a `WaitGroup`. Instead, it uses `time.Sleep` to give goroutines time to finish. This approach is unreliable because the main function may exit before all goroutines complete.


```go
for i := 1; i <= 10000; i++ {
    go func(id int) {
        fmt.Printf("Goroutine %d finished\n", id)
    }(i)
}
// Sleep to give goroutines time to finish (not reliable)
time.Sleep(10 * time.Millisecond)
fmt.Println("Main function finished")
```

**Output:**
```
Goroutine 1 finished
Goroutine 2 finished
...
Main function finished
```
With a large number of goroutines and a short sleep, many goroutines may not finish before the main function exits. This demonstrates why using `time.Sleep` is unreliable for synchronization.

## Summary
- Use `sync.WaitGroup` for reliable goroutine synchronization.
- Using `time.Sleep` is not recommended for waiting on goroutines.

## Learn More
- [sync.WaitGroup documentation](https://pkg.go.dev/sync#WaitGroup)
- [Go Concurrency Patterns](https://golang.org/doc/effective_go#concurrency)
