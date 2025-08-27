# Livelock Example in Go

This example demonstrates a livelock scenario using two goroutines and two locks in Go.

## What is Livelock?

A livelock occurs when two or more threads (or goroutines) are actively trying to avoid deadlock by releasing and reacquiring resources, but as a result, they never make progress. Unlike deadlock, the program is not stuck; it is actively running, but the tasks are not completing.

## Code Overview

```go
// ...existing code...
type stick struct{ sync.Mutex }

func main() {
    var left, right stick
    done := make(chan struct{})

    eat := func(name string, first, second *stick) {
        for i := 0; i < 10; i++ {
            first.Lock()
            time.Sleep(1 * time.Millisecond) // Try to avoid deadlock
            if !second.TryLock() {
                first.Unlock()
                fmt.Printf("%s: couldn't acquire both sticks, retrying\n", name)
                time.Sleep(1 * time.Millisecond)
                continue // Livelock: both keep retrying
            }
            fmt.Printf("%s: eating\n", name)
            second.Unlock()
            first.Unlock()
            time.Sleep(1 * time.Millisecond)
        }
        done <- struct{}{}
    }

    go eat("A", &left, &right)
    go eat("B", &right, &left)
    <-done
    <-done
}
```

## How Livelock Happens Here

- Two goroutines (A and B) try to acquire two locks (sticks) in opposite order.
- If a goroutine can't acquire both locks, it releases the first and retries.
- Both goroutines may keep releasing and retrying, so neither can eat, even though they are both active.
- This is livelock: the program is running, but not making progress.

## How to Run

From the `4. livelock` directory, run:
```sh
go run main.go
```
You will see output like:
```
A: couldn't acquire both sticks, retrying
B: couldn't acquire both sticks, retrying
... (repeats)
```
Sometimes, one goroutine may succeed and print "eating", but often both will keep retrying.

## Learn More
- [Livelock (Wikipedia)](https://en.wikipedia.org/wiki/Livelock)
- [Go sync package](https://pkg.go.dev/sync)
