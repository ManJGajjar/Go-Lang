package main 

import "fmt"

func pointerExample() {
	var a *int
	b := 10
	a = &b	

	fmt.Println("Value of b:", b)
	fmt.Println("Address of b:", &b)
	fmt.Println("Value of a (address of b):", a)
	fmt.Println("Value pointed to by a:", *a)
}