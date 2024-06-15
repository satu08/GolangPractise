package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // lock the variable to access by other goroutine
			v := counter
			runtime.Gosched() // yields another go routine to run
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("go routine- ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("number of cpu's ", runtime.NumCPU())
	fmt.Println("goroutine ", runtime.NumGoroutine())
	fmt.Println("counter : ", counter)
}
