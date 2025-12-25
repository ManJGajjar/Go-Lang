package main

import "fmt"

func deferExample() {
	defer fmt.Println("This will be printed at the end of the function execution")
	fmt.Println("This will be printed first")
}

func gotoExample() {
	fmt.Println("Demonstrating goto statement")
	var i int = 0
loop:
	fmt.Println(i)
	i++
	if i < 5 {
		goto loop
	}
}	