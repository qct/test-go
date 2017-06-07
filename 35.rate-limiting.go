package main

import (
    "time"
    "fmt"
)

func main() {
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(time.Millisecond * 500)

    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }
}
