package main

import "fmt"

func main() {
	// send only channel
	ch := make(chan<- int, 2)
	ch <- 42
	ch <- 43
	//	fmt.Println(<-ch)
	//	fmt.Println(<-ch)
	fmt.Println("-------------")
	fmt.Printf("%T\n", ch)

	// receive only channel
	ch2 := make(<-chan int, 2)
	//	ch2 <- 42
	//	ch2 <- 43
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println("-------------")
	fmt.Printf("%T\n", ch)

}
