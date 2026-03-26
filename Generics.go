package main

import "fmt"

type Stringer interface {
   String() string
}
func PrintString[T Stringer](value T) {
   fmt.Println(value.String())
}
type MyType struct {
   data string
}
func (m MyType) String() string {
   return m.data
}
func main() {
   val := MyType{data: "Hello, Generics!"}
   PrintString(val) 
}