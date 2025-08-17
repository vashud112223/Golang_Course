package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating map
	// m := make(map[string]string)

	// setting an element
	// m["course"] = "golang"
	// m["area"] = "backend"

	// get an element
	// fmt.Println(m["course"], m["area"])

	// if key doesn't exists in the map then it returns zero
	// fmt.Println(m["hell"])
	// delete(m, "area")
	// clear(m)
	// fmt.Println(m)

	// m := map[string]int{"price": 40, "phone": 3}

	// val, key := m["phone"]

	// if key {
	// 	fmt.Println("present", val)
	// } else {
	// 	fmt.Println("not present")
	// }
	// fmt.Println(m)

	m1 := map[string]int{"phone": 4, "price": 1000}
	m2 := map[string]int{"phone": 4, "price": 1000}

	fmt.Println(maps.Equal(m1, m2))

}
