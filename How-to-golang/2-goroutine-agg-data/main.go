package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	fmt.Println("aggregating data using goroutine")

	username := fetchUser()
	reschn := make(chan any, 2)

	wg.Add(2)
	go fetchUserLikes(username, reschn)
	go fetchMatch(username, reschn)
	wg.Wait()
	close(reschn)
	for res := range reschn {
		fmt.Println("response: ", res)
	}
	fmt.Println("elapsed:", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "satya"
}

func fetchUserLikes(userName string, reschn chan any) {
	time.Sleep(time.Millisecond * 100)
	reschn <- 87
	wg.Done()
}

func fetchMatch(userName string, reschn chan any) {
	time.Sleep(time.Millisecond * 100)
	reschn <- "shraddha"
	wg.Done()
}
