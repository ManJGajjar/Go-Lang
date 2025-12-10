//Impementaion of Go variables 
package main

import "fmt"

func PrintVariablesExample() {
	// Declaring variables
	var a int = 10
	var b string = "GoLang"
	var c bool = true
	var d float64 = 20.5
	e := 19 // Short declaration

	// Printing variables
	fmt.Println("Integer:", a)
	fmt.Println("String:", b)
	fmt.Println("Boolean:", c)
	fmt.Println("Float:", d)
	fmt.Println("Short Declared Float:", e)
}