package main

import "fmt"

type Rectangle struct {
    Width  int
    Length int
}

// Struct receiver method to calculate area
func (r Rectangle) Area() int {
    return r.Width * r.Length
}

// Pointer receiver method to scale dimensions
func (r *Rectangle) Scale(factor int) {
    r.Width *= factor
    r.Length *= factor
}

func MethodsEx() {
    // Initialize a Rectangle
    rect := Rectangle{Width: 5, Length: 3}

    // Call the Area method (struct receiver)
    fmt.Println("Initial Area:", rect.Area())

    // Call the Scale method (pointer receiver)
    rect.Scale(10)

    // Check updated area and dimensions
    fmt.Println("Scaled Area:", rect.Area())
    fmt.Println("Updated Width:", rect.Width)
    fmt.Println("Updated Length:", rect.Length)
}