# Go Concurrency Examples

This repository contains examples demonstrating Go's concurrency features.

## Structure

- `routines/` - Basic goroutine examples

## Running the Examples

### Basic Goroutines

```bash
cd routines
go run main.go
```

This example demonstrates:
- Creating multiple goroutines using the `go` keyword
- Basic synchronization using `time.Sleep`
- Concurrent execution of worker functions

## Prerequisites

- Go 1.x or later

## Output Example

```
Starting goroutines...
Worker 1 starting
Worker 2 starting
Worker 3 starting
Worker 4 starting
Worker 5 starting
Worker 1 done
Worker 2 done
Worker 3 done
Worker 4 done
Worker 5 done
All workers completed
```