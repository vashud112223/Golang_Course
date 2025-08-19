package main

import "fmt"

func sum(num *int) {
	*num = 5
	fmt.Println("In function", *num)
}
func main() {
	num := 1
	sum(&num)
	fmt.Println("After changed", num)
}
