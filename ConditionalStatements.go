package main

import "fmt"

func ConditionalStatements() {
	var a int = 10
	var b int = 20

	fmt.Println("Exmaple for Basic If Statment ")
	if a < b {
		fmt.Println("a is greater than b")
	}

	fmt.Println("Example for If-Else Statement")
	if a%2 == 0 {
		fmt.Println("a is even")
	} else {
		fmt.Println("a is odd")
	}

	fmt.Println("Example for if-else if-else Statement")
	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}
}