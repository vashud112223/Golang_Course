package main

import (
	"fmt"
	"sync"
)

// Mutex -> to avoid race condition

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(w *sync.WaitGroup) {
	defer func() {
		w.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.views += 1

}
func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println("Views:", myPost.views)
}
