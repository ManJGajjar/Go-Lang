package main

import "fmt"
 
var g int = 20
 
func VariableScope() {
   var g int = 10
 
   fmt.Printf ("value of g = %d\n",  g)
}