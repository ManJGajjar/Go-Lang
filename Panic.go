package main

import "fmt"

func divide(a, b int) int {
    if b == 0 {
        panic("Division by zero!")
    }
    return a / b
}

func PanicEx() {
    fmt.Println(divide(4, 2)) 
    fmt.Println(divide(4, 0)) 
    fmt.Println(divide(8, 2)) 
}