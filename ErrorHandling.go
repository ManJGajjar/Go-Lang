package main

import (
    "errors"
    "fmt"
)

func checkNumber(num int) (string, error) {
    if num < 0 {
        return "", errors.New("number is negative")
    }
    return "number is positive", nil
}

func EHExample() {
    result, err := checkNumber(-5)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println(result)
    }
}