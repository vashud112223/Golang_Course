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

	product := "WCE"

	switch product {
	case "Tractor":
		fmt.Println("product is Tractor")
	case "WC":
		fmt.Println("product is WC")
	case "HL":
		fmt.Println("product is HL")
	case "Nano Loan":
		fmt.Println("product is Nano Loan")
	default:
		fmt.Println(" no product found")
	}
}
