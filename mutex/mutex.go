package main

import (
	"fmt"
	"sync"
)

type post struct {
	view int

	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()

	p.mu.Lock()
	p.view += 1

	p.mu.Unlock()
}

func main() {

	var wg sync.WaitGroup
	myPost := post{view: 0}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go myPost.inc(&wg)

	}
	fmt.Println(myPost.view)

}
