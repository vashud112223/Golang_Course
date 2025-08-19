package main

import "fmt"

// higher order functions
func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}
func main() {

	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
}
