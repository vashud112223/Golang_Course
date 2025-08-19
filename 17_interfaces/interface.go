package main

import "fmt"

//interface -> acts like contract
type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	// gateway stripe
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// razorPaymentGw.pay(amount)
	// stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("make payment using razorpay", amount)
}

func (r razorpay) refund(amount float32, account string) {
	fmt.Println("make payment using razorpay", amount)
}

type stripe struct{}

func (r stripe) pay(amount float32) {
	fmt.Println("make payment using stripe", amount)
}

func main() {
	// newPayment := payment{}
	// newPayment.makePayment(100)
	// stripePaymentGw := stripe{}
	// newPayment := payment{
	// 	gateway: stripePaymentGw,
	// }
	// newPayment.makePayment(100)

	razorPaymentGw := razorpay{}
	newPayment1 := payment{
		gateway: razorPaymentGw,
	}
	newPayment1.makePayment(100)
}
