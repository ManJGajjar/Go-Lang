package main

import "fmt"

type Stringer[T any] interface {
	String() string
}

type MyType struct {
	data string
}
func (m MyType) String() string {
	return m.data
}
func PrintString[T Stringer[T]](value T) {
	fmt.Println(value.String())
}

func GenericInterfacesEx() {
	val := MyType{data: "Hello, Peter Griffin!"}
	PrintString(val)
}