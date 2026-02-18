package main

import "fmt"

func sender(ch chan int) {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch) lues
}

func ClosingChannelEx() {
    ch := make(chan int)

    go sender(ch)

    for val := range ch {
        fmt.Println(val)
    }

    fmt.Println("Channel is closed.")
}