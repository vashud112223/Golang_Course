package main

import "fmt"

// func printSlice(items []int) {
// 	fmt.Println(items)
// }

// func printstringSlice(items []string) {
// 	fmt.Println(items)
// }

// We need generics to use the same code for different types
// func genericSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// wecan use this also
func genericSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

type stack[T any] struct {
	element []T
}

func main() {
	// nums := []int{1, 2, 3}
	// printSlice(nums)

	nums := []bool{true, false, true}
	// printstringSlice(nums)
	genericSlice(nums, "Ashu")

	mystack := stack[string]{
		element: []string{"go", "java", "javascript"},
	}

	fmt.Println(mystack)
}
