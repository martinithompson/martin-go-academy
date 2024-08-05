// main.go
package main

import (
	"log"
	"net/http"
)

const port = ":5000"

func main() {
	server := NewTodoServer(NewInMemoryTodoStore())
	log.Printf("Starting api on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, server))
}
