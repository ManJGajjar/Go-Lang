package main
  
import "fmt"
  
func MiscOperators() {
  a := 4
   
  b := &a 
  fmt.Println(*b) 
  *b = 7 
  fmt.Println(a) 
}