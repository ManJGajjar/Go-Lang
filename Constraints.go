package main

import (
   "fmt"
   "golang.org/x/exp/constraints"
)

func Add[T constraints.Integer | constraints.Float](a, b T) T {
   return a + b
}
func ConstraintsEx() {
   fmt.Println(Add(1, 2))
   fmt.Println(Add(1.5, 2.5))
}