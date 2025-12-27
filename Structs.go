package main

import "fmt"

func StructExample() {
	type Person struct {
		Name   string
		Age    int
		Gender string
	}

	p1 := Person{Name: "Peter", Age: 24, Gender: "Male"}
	fmt.Println("Person Details:")
	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)
	fmt.Println("Gender:", p1.Gender)
}