package main

import "fmt"

func divide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
            result = 0 // Set result to 0 or any default value after recovery
        }
    }()
    if b == 0 {
        panic("Division by zero!")
    }
    return a / b
}

func RecoverEx() {
    fmt.Println(divide(4, 2)) // Executes successfully
    fmt.Println(divide(4, 0)) // Recovers from panic
    fmt.Println(divide(8, 2)) // Continues execution
}