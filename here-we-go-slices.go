package main

import "fmt"

func main() {
	// Declare it similar to an array
	numSlice := []int{5, 4, 3, 2, 1 }
	// But you can slice from another slice into a new slice
	// Notation is start_index:end_index
	numSlice2 := numSlice[3:5]
	fmt.Println("numSlice2[0] = ", numSlice2[0])
	fmt.Println("numSlice[:2] = ", numSlice[:2])
	fmt.Println("numSlice[2:] = ", numSlice[2:])

	// Here the slice starts off with 5 values of int to a maximum of 10
	numSlice3 := make([]int, 5, 10)
	copy(numSlice3, numSlice)
	for _, value := range numSlice3 {
		fmt.Println("value = ", value)
	}

	fmt.Println("\n")

	numSlice3 = append(numSlice3, 0, -1)
	for _, value := range numSlice3 {
		fmt.Println("value = ", value)
	}
}
