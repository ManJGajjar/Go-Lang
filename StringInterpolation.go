package main

import "fmt"    

func StrInter() {
   name := "Ritika"  
   age := 21    
   fmt.Println("The string interpolation can be demonstrated as:")
   message := fmt.Sprintf("Hello, my name is %s and I am %d years old", name, age) 
   fmt.Println(message) 
}