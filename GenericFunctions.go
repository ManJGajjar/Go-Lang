package main

import "fmt"

func Print[T any](value T) {
   fmt.Println(value)
}
func GenFuncEx() {
   Print(42)       
   Print("Hello")   
   Print(3.14)      
}