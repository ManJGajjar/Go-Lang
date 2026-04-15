package main

import "fmt"

type Container[T any] struct {
	Value T
}

type Box[T any] struct {
	Container[T]
}

func main() {
	box := Box[int]{Container[int]{Value: 64}}
	fmt.Println("box.Value:", box.Value) 
}