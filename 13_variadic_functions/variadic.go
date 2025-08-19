package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total

}

// A variadic function lets you pass zero or more arguments of the same type. Internally, Go treats these arguments as a slice.

func main() {
	nums := []int{3, 4, 5, 6}
	result := sum(nums...)
	fmt.Println(result)
}
