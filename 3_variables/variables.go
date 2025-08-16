package main

// fmt-> formatter
//println-> print line
import "fmt"

func main() {
	// we need this syntax when we use this value later
	var name string = "Golang"
	// we can also use it
	var s = "hello"
	// shorthand -> we can use it when we declaring the value at the time of initialization
	val := "world"

	var my string
	my = "learning"

	int := 32
	fmt.Println(name)
	fmt.Println(s)
	fmt.Println(val)
	fmt.Println(my)
	fmt.Println(int)
}
