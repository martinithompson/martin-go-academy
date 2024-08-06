package main

import (
	"log"
	"net/http"
	"strconv"
)

const port = 8000

func main() {
	server := Server{}
	server.cmds = startTodoManager()

	http.HandleFunc("/todos", server.todosHandler)
	http.HandleFunc("/todos/", server.todoHandler)

	log.Printf("API server starting on port %d\n", port)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(port), nil))
}
