

# Go Concurrency Examples

**Repository Description:**

This repository is a comprehensive collection of runnable, well-documented examples covering all major concurrency primitives and patterns in Go. It is designed for learners, educators, and developers who want to understand, teach, or demonstrate Go's approach to concurrent programming through clear, minimal, and practical code samples.

---

## About

This repository contains practical, self-contained examples of Go concurrency primitives and patterns. Each folder demonstrates a specific concept with runnable code and explanations. The goal is to help you understand how to use goroutines, channels, mutexes, and other synchronization tools in real-world scenarios, as well as to illustrate common pitfalls such as race conditions, deadlocks, and goroutine leaks.


## Contents

- **1. routines/** — Basic goroutine usage
- **2. race-condition/** — Race condition and its fix
- **3. deadlock/** — Deadlock scenario
- **4. livelock/** — Livelock scenario
- **5. starvation/** — Starvation scenario
- **6. wait-group/** — Using and not using sync.WaitGroup
- **7. mutex/** — Mutex, RWMutex, and what happens without them
- **8. cond/** — sync.Cond for goroutine coordination
- **9. once/** — sync.Once for one-time initialization
- **10. channels/** — Channels, buffered channels, select statement
- **11. routine-leaks/** — Goroutine leaks and how to fix them

---

## How to Run

Navigate to any example directory and run the main file:

```sh
go run main.go
```
or for subfolders:
```sh
go run with-wait-group/main.go
go run without-mutex/main.go
go run rw-mutex/main.go
```

---

## Example Overview

### 1. Goroutines
Shows how to launch multiple goroutines and synchronize with `time.Sleep`.

### 2. Race Condition
Demonstrates a race condition when incrementing a shared variable from multiple goroutines, and how to fix it with a mutex and WaitGroup.

### 3. Deadlock
Shows a classic deadlock with two mutexes and two goroutines locking in opposite order.

### 4. Livelock
Demonstrates livelock, where goroutines are active but unable to make progress due to repeated retries.

### 5. Starvation
Shows how a greedy goroutine can starve a polite one by monopolizing a lock.

### 6. WaitGroup
Compares reliable goroutine synchronization with WaitGroup vs. unreliable time.Sleep.

### 7. Mutex & RWMutex
Shows safe and unsafe concurrent access using Mutex and RWMutex, and the problems without them.


### 8. sync.Cond
Demonstrates goroutine coordination using sync.Cond for signaling and waiting on conditions.

### 9. sync.Once
Shows how to use sync.Once to ensure a function is executed only once, even with multiple goroutines.

### 10. Channels
Demonstrates basic channels, buffered channels (including overflow and multiple senders), and the select statement for handling multiple channels.

### 11. Routine Leaks
Shows how goroutine leaks can occur and how to fix them by properly closing channels and signaling completion.

---

## Prerequisites

- Go 1.x or later

---

## Learn More
- [Go Concurrency Patterns](https://golang.org/doc/effective_go#concurrency)
- [sync package documentation](https://pkg.go.dev/sync)