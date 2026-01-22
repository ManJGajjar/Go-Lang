package main

import "fmt"

func check(code int) (string, error) {
    switch code {
    case 400:
        return "", &CustomError{Code: 400, Message: "Invalid Request"}
    case 404:
        return "", &CustomError{Code: 404, Message: "Resource Not Found"}
    default:
        return "Success", nil
    }
}