package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from Goroutine!")
}

func GORoutinesEx() {
    go sayHello()
    time.Sleep(1 * time.Second)
    fmt.Println("Main function finished")
}