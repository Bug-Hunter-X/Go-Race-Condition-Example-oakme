```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var m sync.Mutex
        
        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        m.Lock()
                        fmt.Printf("Goroutine %d locked the mutex\n", i)
                        // Simulate some work
                        // time.Sleep(time.Second)
                        m.Unlock()
                        fmt.Printf("Goroutine %d unlocked the mutex\n", i)
                }(i)
        }
        wg.Wait()
}
```