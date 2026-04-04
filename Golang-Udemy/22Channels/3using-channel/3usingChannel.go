package main

import "fmt"

func main() {

	ch := make(chan int)

	go foo(ch)

	bar(ch)

	fmt.Println("about to exit")
}

func foo(ch chan<- int) {
	ch <- 1
}

func bar(ch <-chan int) {
	fmt.Println(<-ch)
}
