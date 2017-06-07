package main

import (
    "fmt"
    "strconv"
)

func main() {
    messages := make(chan string, 2)

    for i := 0; i < 2; i++ {
        msg := "buffered channel " + strconv.Itoa(i)
        messages <- msg
    }

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
