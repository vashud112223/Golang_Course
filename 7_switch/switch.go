package main

import (
	"fmt"
)

func main() {
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its weekend")
	// default:
	// 	fmt.Println("its weekday")

	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Its an integer")

		case string:
			fmt.Println("Its an string")

		case bool:
			fmt.Println("Its an bool")

		default:
			fmt.Println("other")
		}

	}
	whoAmI(true)

}
