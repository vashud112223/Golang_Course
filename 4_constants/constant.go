package main

import "fmt"

const age = 30

func main() {
	const name = "Golang"
	// name = "java"
	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(host, port)
}
