package main

import "fmt"

// for -> only construct in go for looping
func main() {
	//while loop using for loop
	// i := 1
	// for i < 3 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// // infinite loop
	// for{
	// 	println("1")
	// }

	//classic for loop

	// for i := 0; i < 3; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)

	// }

	//range

	for i := range 11 {
		fmt.Println(i)
	}

}
