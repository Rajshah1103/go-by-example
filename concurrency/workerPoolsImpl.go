package concurrency

import (
	"fmt"
	"time"
)

// example of workpools implementation using go routines and channels

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		// Simulate doing some work
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2 // example of processing the job
	}
}

func InitializeWorkerPool() {
	// This function would typically initialize worker pools
	// and set up channels for communication between workers.
	// For this example, we will just define a constant for the number of jobs.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
