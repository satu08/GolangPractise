package main

import "fmt"

func main() {

	ch := make(chan int)

	go send(ch)
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("about to exit")
}

func send(ch chan<- int) {
	for i := 0; i < 50; i++ {
		ch <- i
	}
	close(ch) // if we do not close this channel then receiver will wait for more values from that channel

}
