// main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewTodoServer(NewInMemoryTodoStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
