package main

import (
    "time"
    "fmt"
)

func main() {
    timeout := make(chan bool, 1)

    go func() {
        time.Sleep(1e9)
        timeout <- true
    }()

    ch := make(chan bool)
    select {
    case <- ch:

    case msg2 := <-timeout:
        fmt.Println(msg2)
    }
}
