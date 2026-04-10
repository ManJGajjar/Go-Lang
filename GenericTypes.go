package main
import "fmt"

type Stack[T any] struct {
   items []T
}
func (s *Stack[T]) Push(item T) {
   s.items = append(s.items, item)
}
func (s *Stack[T]) Pop() T {
   if len(s.items) == 0 {
      panic("stack is empty")
   }
   item := s.items[len(s.items)-1]
   s.items = s.items[:len(s.items)-1]
   return item
}
func main() {
   intStack := Stack[int]{}
   intStack.Push(1)
   intStack.Push(2)
   fmt.Println(intStack.Pop()) 

   stringStack := Stack[string]{}
   stringStack.Push("Go")
   stringStack.Push("Generics")
   fmt.Println(stringStack.Pop())
}