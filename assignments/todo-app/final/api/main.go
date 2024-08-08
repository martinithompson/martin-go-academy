package main

import (
	"log"
	"net/http"
)

const port = "8000"

func main() {
	server := Server{}
	server.cmds = startTodoManager()

	http.HandleFunc("/todos", server.todosHandler)
	// had to create a separate handler for dynamic routes e.g. /todos/{id}
	// not sure if necessary but was only way I could get it to work!
	http.HandleFunc("/todos/", server.todoHandler)

	log.Printf("API server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

// WIP! - add a worker pool
// func main() {

// 	jobs := make(chan int, 10)
// 	results := make(chan int, 10)

// 	// 3 workers
// 	for x := 1; x <= 3; x++ {
// 		go worker(x, jobs, results)
// 	}

// 	// give them jobs
// 	for j := 1; j <= 6; j++ {
// 		jobs <- j
// 	}

// 	close(jobs)

// 	// get the results
// 	for r := 1; r <= 6; r++ {
// 		fmt.Println("result received from worker: ", <-results)
// 	}
// }

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for job := range jobs {
// 		fmt.Println("Worker ", id, " is working on job ", job)
// 		duration := time.Duration(2 * time.Second)
// 		time.Sleep(duration)
// 		fmt.Println("Worker ", id, " completed job ", job)
// 		results <- id
// 	}
// }
