package main

import "fmt"

// range-> iterating over data structures
func main() {
	// nums := []int{6, 7, 8}

	// sum := 0
	// for i, num := range nums {

	// 	sum = sum + num
	// 	fmt.Println(num, i)
	// }
	// fmt.Println(sum)

	// m := map[string]int{"phone": 4, "price": 1000}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// unicode code point rune
	// starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
