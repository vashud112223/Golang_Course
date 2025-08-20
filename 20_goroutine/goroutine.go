package main

import (
	"fmt"
	"sync"
)

// we are using the wait group to execute the goroutine function concurently
func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("daily routine", id)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i < 11; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
	// time.Sleep(2 * time.Second)
}
