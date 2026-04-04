package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	sem <- struct{}{} // acquire
	fmt.Println("Start task", id)

	time.Sleep(1 * time.Second)

	fmt.Println("End task", id)
	<-sem // release
}

func main() {
	const limit = 3
	sem := make(chan struct{}, limit)

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go task(i, sem, &wg)
	}

	wg.Wait()
}
