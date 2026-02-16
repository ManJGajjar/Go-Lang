package main

import (
    "fmt"
    "time"
)

func worker(id int, ch chan int) {
    fmt.Printf("Worker %d: Sending data...\n", id)
    ch <- id
    fmt.Printf("Worker %d: Sent data\n", id)
}

func ChannelSync() {
    ch := make(chan int)

    for i := 1; i <= 3; i++ {
        go worker(i, ch)
    }

    for i := 1; i <= 3; i++ {
        val := <-ch
        fmt.Printf("Main: Received data %d from channel\n", val)
    }

    time.Sleep(2 * time.Second)
}