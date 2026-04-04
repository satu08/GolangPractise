package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go sendchannel(eve, odd, quit)
	receivechannel(eve, odd, quit)
	fmt.Println("about to exit")
}

func sendchannel(eve, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receivechannel(eve, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case i := <-eve:
			fmt.Println("value from even channel ", i)
		case i := <-odd:
			fmt.Println("value from odd channel ", i)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from ok idiom", i, ok)
				return
			} else {
				fmt.Println("value from quit channel ", i)
			}
		}
	}
}
