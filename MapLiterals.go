package main

import "fmt"

func MapLiterals() {
    m := map[string]int{
        "Lewis Hamilton": 44,
        "Max Verstappen": 3,
        "George Russell": 63,
    }

    // Accessing values
    fmt.Println("Lewis Hamilton's number:", m["Lewis Hamilton"]) 	
    fmt.Println("Max Verstappen's number:", m["Max Verstappen"]) 
    fmt.Println("George Russell's number:", m["George Russell"]) 	
}