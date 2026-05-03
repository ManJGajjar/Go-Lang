package main
import "fmt"
func AnonymousExample() {
   greet := func() {
      fmt.Println("Hello, Stewie Griffin.")
   }
   greet() 

   add := func(a, b int) int {
      return a + b
   }
   sum := add(9, 5)
   fmt.Println("Sum =", sum)
}
