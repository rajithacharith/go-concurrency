# Mutex and RWMutex Examples in Go

This directory demonstrates the use of `sync.Mutex` and `sync.RWMutex` for safe concurrent access in Go, and shows what happens when you do not use them.

## Structure

- `with-mutex/main.go`: Uses a mutex to safely increment a shared counter.
- `without-mutex/main.go`: No mutex; increments a shared counter unsafely (race condition).
- `rw-mutex/main.go`: Uses a read-write mutex to allow multiple readers and one writer.
- `without-rw-mutex/main.go`: No mutex; concurrent reads and writes to a shared value (race condition).

---

## What is a Mutex?
A mutex (mutual exclusion lock) is used to protect shared data from being accessed by multiple goroutines at the same time. Only one goroutine can hold the lock at a time.

**Example:**
```go
var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()
```

- In `with-mutex/main.go`, 1000 goroutines increment a shared counter using a mutex. The final value is always correct.
- In `without-mutex/main.go`, 1000 goroutines increment a shared counter without a mutex. The final value is usually wrong due to race conditions.

---

## What is a RWMutex?
A read-write mutex allows multiple readers or one writer at a time. Use `RLock` for reading and `Lock` for writing.

**Example:**
```go
var rw sync.RWMutex
rw.RLock()   // for reading
// read shared data
rw.RUnlock()
rw.Lock()    // for writing
// write shared data
rw.Unlock()
```

- In `rw-mutex/main.go`, one writer updates a value while multiple readers read it safely using `RWMutex`.
- In `without-rw-mutex/main.go`, readers and writers access the value concurrently without synchronization, leading to race conditions and inconsistent results.

---

## How to Run

From the `7. mutex` directory, run any example:
```sh
go run with-mutex/main.go
go run without-mutex/main.go
go run rw-mutex/main.go
go run without-rw-mutex/main.go
```

## Summary
- Use `sync.Mutex` for safe exclusive access to shared data.
- Use `sync.RWMutex` for safe concurrent reads and exclusive writes.
- Not using these locks can lead to race conditions and unpredictable results.

## Learn More
- [sync.Mutex documentation](https://pkg.go.dev/sync#Mutex)
- [sync.RWMutex documentation](https://pkg.go.dev/sync#RWMutex)
- [Go Concurrency Patterns](https://golang.org/doc/effective_go#concurrency)
