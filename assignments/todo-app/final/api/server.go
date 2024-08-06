package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	todos "todo-app/project/todos"
)

type Server struct {
	cmds chan<- Command
}

func (s *Server) todosHandler(w http.ResponseWriter, r *http.Request) {
	replyChan := make(chan []todos.Todo)
	switch r.Method {
	case http.MethodPost:
		log.Printf("POST %v", r)
		var todo todos.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil || todo.Name == "" {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		s.cmds <- Command{ct: AddCommand, todoItem: todo, replyChan: replyChan}
		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		log.Printf("GET %v", r)
		s.cmds <- Command{ct: GetCommand, replyChan: replyChan}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	reply := <-replyChan

	// return an empty array if no todos have been added
	if reply == nil {
		reply = []todos.Todo{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)
}

func (s *Server) todoHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	replyChan := make(chan []todos.Todo)

	switch r.Method {
	case http.MethodPatch:
		log.Printf("PATCH %v", r)
		var todo todos.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil || todo.Name == "" {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if id != todo.Id {
			http.Error(w, "Id mismatch", http.StatusBadRequest)
			return
		}
		s.cmds <- Command{ct: EditCommand, todoItem: todo, replyChan: replyChan}
		w.WriteHeader(http.StatusCreated)
	case http.MethodDelete:
		log.Printf("DELETE %v", r)
		s.cmds <- Command{ct: DeleteCommand, todoItem: todos.Todo{Id: id}, replyChan: replyChan}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	reply := <-replyChan

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)
}
