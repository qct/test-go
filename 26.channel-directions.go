package main

import "fmt"

func ping(pings chan <- string, msg string) {
    fmt.Println("write msg to pings")
    pings <- msg
}

func pong(pings <-chan string, pongs chan <- string) {
    msg := <-pings
    fmt.Println("transfer msg from pings to pongs")
    pongs <- msg
}
func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
