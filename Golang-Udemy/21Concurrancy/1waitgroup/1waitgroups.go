package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// The Main function is first go routine
	fmt.Println("Introduction to concurrency")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("------------------")
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Wait() // wait for wg.Done function in foo() and bar()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar ", i)
	}
	wg.Done()
}
