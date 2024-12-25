# Go Race Condition Example

This repository contains a simple Go program that demonstrates a race condition.  Race conditions occur when multiple goroutines access and modify shared resources concurrently without proper synchronization mechanisms. This can lead to unexpected and inconsistent results.

The example uses a mutex to protect a shared resource (in this case, just printing to the console).  However, the way the goroutines are launched,  it's possible for them to contend for the mutex in unpredictable ways.  

## Running the Example

1. Clone this repository.
2. Navigate to the repository directory.
3. Run the program using `go run bug.go`.
4. Observe the output. You'll notice that the order of mutex locks and unlocks might be inconsistent across different runs, which could lead to unexpected behavior in real-world scenarios, such as data corruption or program crashes.

## Solution

The solution provided in `bugSolution.go` demonstrates how to correct this issue.