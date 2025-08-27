# Go Concurrency Race Condition Example

This project demonstrates a race condition in Go and how to fix it using synchronization primitives.

## Structure

```
race-condition/
  original/
    main.go   # Race condition example
  fixed/
    main.go   # Fixed example using sync.Mutex and sync.WaitGroup
```

## How to Run

Navigate to the `race-condition` directory and run:

### Race Condition Example
```
go run original/main.go
```
**Expected Output:**
The final counter value will likely be less than 1000 due to concurrent access to the shared variable `counter` without synchronization. This demonstrates a race condition.

### Fixed Example
```
go run fixed/main.go
```
**Expected Output:**
The final counter value will always be 1000. This is because the code uses a mutex to synchronize access to the shared variable and a wait group to wait for all goroutines to finish.

## Explanation

- **Race Condition Example:**
  - 1000 goroutines increment a shared variable `counter`.
  - No synchronization is used, so multiple goroutines may update `counter` at the same time, causing lost updates.
  - The result is usually less than 1000.

- **Fixed Example:**
  - Uses `sync.Mutex` to lock access to `counter` so only one goroutine can update it at a time.
  - Uses `sync.WaitGroup` to wait for all goroutines to finish before printing the result.
  - The result is always 1000, showing correct synchronization.

## Learn More
- [Go Concurrency](https://golang.org/doc/effective_go#concurrency)
- [Race Conditions](https://en.wikipedia.org/wiki/Race_condition)
- [Go sync package](https://pkg.go.dev/sync)
