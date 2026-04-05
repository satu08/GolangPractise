package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter:", atomic.LoadInt64(&counter))
			runtime.Gosched() // yields another go routine to run
			wg.Done()
		}()
		fmt.Println("go routine- ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("number of cpu's ", runtime.NumCPU())
	fmt.Println("goroutine ", runtime.NumGoroutine())
	fmt.Println("counter : ", counter)
}
