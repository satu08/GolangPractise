package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	go busyLoop()
	go busyLoop()
	fmt.Println("satyanarayan")
	fmt.Println(runtime.GC)

	done := make(chan struct{})

	go func() {
		fmt.Println("work done")
		done <- struct{}{}
	}()
	<-done
	fmt.Println("satya")

	// unbuffred channel
	jobs := make(chan int, 10)

	for i := 0; i < 3; i++ {
		go worker(jobs)
	}

	for j := 0; j < 100; j++ {
		jobs <- j
	}

	select {}

}

func busyLoop() {
	//time.Sleep(time.Second * 2)
	fmt.Println("satyanarayan jadhav")
}

func worker(job chan int) {
	for v := range job {
		fmt.Println(v)
	}
}
