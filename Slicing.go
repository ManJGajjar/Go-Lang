package main

import ("fmt")

func Slicing() {
	//Creating Slices
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[1:4] // elements 2,3,4
	fmt.Println("From array:", s1)

	// Using make
	s2 := make([]int, 3, 5) // len=3, cap=5
	fmt.Println("Using make:", s2, "len:", len(s2), "cap:", cap(s2))

	// Slice literal
	s3 := []int{10, 20, 30}
	fmt.Println("Slice literal:", s3)

	// Nil slice vs empty slice
	var nilSlice []int
	emptySlice := []int{}
	fmt.Println("Nil slice:", nilSlice, nilSlice == nil)
	fmt.Println("Empty slice:", emptySlice, emptySlice == nil)

	//Slicing Operations
	base := []int{0, 1, 2, 3, 4, 5}

	fmt.Println("base[:3]  ->", base[:3])  // from start
	fmt.Println("base[2:]  ->", base[2:])  // to end
	fmt.Println("base[1:4] ->", base[1:4]) // middle
	fmt.Println("base[:]   ->", base[:])   // full slice

	//Reslicing
	r := base[1:5]
	fmt.Println("Reslice:", r)

	r = r[1:3]
	fmt.Println("Resliced again:", r)

	fs := base[1:3:3]
	fmt.Println("Full slice:", fs, "len:", len(fs), "cap:", cap(fs))

	//Append Method
	a := []int{1, 2}
	a = append(a, 3)
	a = append(a, 4, 5)
	fmt.Println("Append:", a)

	// Append one slice to another
	b := []int{6, 7}
	a = append(a, b...)
	fmt.Println("Append slice:", a)

	//Copy Method
	src := []int{100, 200, 300}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("Copied slice:", dst)

	//Deleting Elements
	del := []int{1, 2, 3, 4, 5}
	i := 2 // index to delete
	del = append(del[:i], del[i+1:]...)
	fmt.Println("After delete:", del)

	//Overlapping Slices
	shared := []int{1, 2, 3, 4}
	x := shared[:2]
	y := shared[1:3]
	x[1] = 99
	fmt.Println("Shared base:", shared)
	fmt.Println("x:", x)
	fmt.Println("y:", y)


	str := "hello"
	fmt.Println("String slice (bytes):", str[1:4])

	// Correct way for Unicode
	runes := []rune("héllo")
	fmt.Println("String slice (runes):", string(runes[1:4]))

	//Slicing Bytes
	fmt.Println("Byte slice:", bytes[1:3])

	sum := addAll([]int{1, 2, 3}...)
	fmt.Println("Variadic slice:", sum)

	//Capacity Growth
	grow := []int{}
	lastCap := cap(grow)
	for i := 0; i < 20; i++ {
		grow = append(grow, i)
		if cap(grow) != lastCap {
			fmt.Println("Capacity grew to:", cap(grow))
			lastCap = cap(grow)
		}
	}
}

func addAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
