package main

import "fmt"

func main() {
	//ch := make(chan int)
	//// this is not going to run because channels block
	//ch <- 42
	//fmt.Println(<-ch)

	//ch := make(chan int)
	//// this will run because it will run in separate goroutine
	//go func() {
	//	ch <- 42
	//}()
	//fmt.Println(<-ch)

	//ch := make(chan int, 1)
	//// this will run we are creating a one buffer
	//ch <- 42
	//fmt.Println(<-ch)

	//ch := make(chan int, 1)
	//// this will not run because we have only one buffer, and we are adding two values in a channel
	//ch <- 42
	//ch <- 43
	//fmt.Println(<-ch)

	ch := make(chan int, 2)
	// this will run because we have two buffers, and we are adding two values in a channel
	ch <- 42
	ch <- 43
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
