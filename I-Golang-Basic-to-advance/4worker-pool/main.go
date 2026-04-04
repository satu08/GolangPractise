package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data int
}

type Result struct {
	JobID int
	Value int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(time.Millisecond * 500)
		results <- Result{JobID: job.ID, Value: job.Data * 2}
	}
}

func main() {
	numWorkers := 3
	numJobs := 10

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- Job{ID: i, Data: i * 10}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
