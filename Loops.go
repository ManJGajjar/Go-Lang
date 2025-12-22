package main

import "fmt"

func loops () {
	// Example of for loop
	fmt.Println("For Loop Example:")
	for i := 1; i <= 5; i++ {

		fmt.Println("Iteration:", i)

	}
	fmt.Println("For as While Loop Example:")
	i := 1
	for i <= 5 {
		fmt.Println("Iteration:", i)
		i++
	}

	/*fmt.Println("Infinite Loop Example")
	for {
		fmt.Println("Hello World !")
	}*/

	fmt.Println("With Range Loop Example:")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}