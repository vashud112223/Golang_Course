package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 12 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Child")
	}

	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission {
		fmt.Println("yes")
	}

	// we can declare variable inside if block
	if age := 20; age >= 18 {
		fmt.Println("person is an adult", age)
	}
	// go does not have terninary operator
}
