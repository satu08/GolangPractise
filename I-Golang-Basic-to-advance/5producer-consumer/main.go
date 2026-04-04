package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Produced:", i)
		ch <- i
	}
	close(ch) // important
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch {
		fmt.Println("Consumed:", val)
	}
}

func main() {
	ch := make(chan int, 2) // buffer = 2
	var wg sync.WaitGroup

	wg.Add(1)
	go consumer(ch, &wg)

	producer(ch)

	wg.Wait()
}
