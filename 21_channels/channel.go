package main

import (
	"fmt"
)

// receiver
// func task(numChan chan int) {
// 	// fmt.Println("Message received", <-numChan)
// 	for num := range numChan {
// 		fmt.Println("Message received", num)
// 		time.Sleep(time.Second)
// 	}
// }

// sender
// func process(numChan chan int, num1 int, num2 int) {
// 	total := num1 + num2
// 	numChan <- total

// }

// goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("processing....")
// }

// func emailSender(emailChan <-chan string, done chan<- bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		chan1 <- "hlo"
	}()
	go func() {
		chan2 <- "Ashu"
	}()
	for i := 0; i < 2; i++ {
		select {
		case chna1Val := <-chan1:
			fmt.Println("received data on chan1", chna1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data on chan2", chan2Val)

		}

	}

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)
	// for i := 0; i < 3; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("done sending")
	// close(emailChan)
	// <-done

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// done := make(chan bool)

	// go task(done)

	// <-done //block

	// numChan := make(chan int) // unbuffered channel

	// go process(numChan, 4, 5)
	// res := <-numChan // blocking
	// fmt.Println(res)
	// numChan := make(chan int)

	// go task(numChan)

	// for {
	// 	numChan <- rand.Intn(100) // blocking
	// }
	// numChan <- "hlo"
	// messageChan := make(chan string)

	// messageChan <- "ping" //blocking

	// msg := <-messageChan

	// fmt.Println(msg)
}
