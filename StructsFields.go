package main

import "fmt"

// Define the Car struct
type Car struct {
    Name  string
    Type  string
    Brand string
    Years int
}

func StructFields() {
    // Initialize a Car instance
    car := Car{
        Name:  "RB19",
        Type:  "FUEL",
        Brand: "RED BULL RACING",
        Years: 1,
    }

    // Accessing fields
    fmt.Println("Car Name:", car.Name)
    fmt.Println("Years in Service:", car.Years)
}