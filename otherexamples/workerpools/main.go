package main

import (
	"fmt"
	"time"
)

func main() {

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// 3 workers
	for x := 1; x <= 3; x++ {
		go worker(x, jobs, results)
	}

	// give them jobs
	for j := 1; j <= 6; j++ {
		jobs <- j
	}

	close(jobs)

	// get the results
	for r := 1; r <= 6; r++ {
		fmt.Println("result received from worker: ", <-results)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker ", id, " is working on job ", job)
		duration := time.Duration(2 * time.Second)
		time.Sleep(duration)
		fmt.Println("Worker ", id, " completed job ", job)
		results <- id
	}
}
