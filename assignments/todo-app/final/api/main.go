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
