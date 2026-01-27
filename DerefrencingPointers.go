package main

import "fmt"

func DRPointersEx() {
   var x int = 64   
   var p *int        
   p = &x          

   fmt.Println("Value of x=", x)
   fmt.Println("Address of x=", &x)
   fmt.Println("Value of p=", p)
   fmt.Println("Dereferenced value of p=", *p)
   *p = 21 
   fmt.Println("New value of x=", x)
}