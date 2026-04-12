package main

import "fmt"

func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func GenericSliceMapsEx() {
	nums := []int{2, 4, 6}
	doubled := Map(nums, func(n int) int {
		return n * 2
	})
	fmt.Println("doubled:", doubled) 
}