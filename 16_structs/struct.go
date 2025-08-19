package main

import (
	"fmt"
	"time"
)

// struct embeding

type customer struct {
	name  string
	phone string
}

// composition, inheritance, code reusablity
type order struct {
	id          string
	amount      float32
	status      string
	createdTime time.Time
	customer
}

// constructor
func newOrder(id string, amount float32, status string) *order {
	//initial setup
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// created methods
// receiver type
func (o *order) changeStatus(status string) {
	o.status = status // no need to dereference it, struct will take care of it
}

func (o order) getAmount() float32 {
	return o.amount // when we are returning , no need to give pointer
}
func main() {
	// instantiate
	myOrder := order{
		id:     "1",
		amount: 30.4,
		status: "pending",
		customer: customer{
			name:  "Ashu",
			phone: "9129189319",
		},
	}
	// myOrder := newOrder("1", 50, "accepted")

	myOrder2 := order{
		id:     "2",
		amount: 50.4,
		status: "received",
	}

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// myOrder.createdTime = time.Now()
	// myOrder.changeStatus("paid")
	fmt.Println(myOrder)
	// fmt.Printf("%+v\n", *myOrder)
	fmt.Println(myOrder2)
	// fmt.Println(myOrder.getAmount())
}
