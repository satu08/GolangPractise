package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

// CPU-heavy function
func heavyWork(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	result := 0.0
	for i := 0; i < 1_000_000; i++ {
		result += math.Sin(float64(i)) * math.Cos(float64(i))
	}

	fmt.Printf("Goroutine %d done, result: %f\n", id, result)
}

func main() {
	// Use all CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create CPU profile file
	f, err := os.Create("cp.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	var wg sync.WaitGroup

	numGoroutines := 50

	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go heavyWork(i, &wg)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("All goroutines finished in:", elapsed)
}
