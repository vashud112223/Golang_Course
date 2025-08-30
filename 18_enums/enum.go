package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeStatus(status OrderStatus) {
	fmt.Println("message status is", status)
}

func main() {
	changeStatus(Delivered)
}
