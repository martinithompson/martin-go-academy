package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	todos "todo-app/project/todos"
)

type Store struct {
	todos todos.Todos
}

func main() {
	server := Server{}
	server.cmds = startTodoManager()

	http.HandleFunc("/todos", server.todosHandler)
	http.HandleFunc("/todos/", server.todoHandler)

	portnum := 8000
	if len(os.Args) > 1 {
		portnum, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("API server starting on port %d\n", portnum)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portnum), nil))
}

type CommandType int

const (
	GetCommand CommandType = iota
	AddCommand
	EditCommand
	DeleteCommand
)

type Command struct {
	ct        CommandType
	todoItem  todos.Todo
	replyChan chan []todos.Todo
}

func startTodoManager() chan<- Command {
	s := Store{}

	cmds := make(chan Command)

	go func() {
		for cmd := range cmds {
			switch cmd.ct {
			case GetCommand:
				fmt.Println("GetCommand")
				cmd.replyChan <- s.todos.Items
			case AddCommand:
				fmt.Println("AddCommand")
				s.todos.AddTaskItems(cmd.todoItem.Item)
				cmd.replyChan <- s.todos.Items
			case EditCommand:
				fmt.Println("EditCommand")
				s.todos.UpdateTodoItem(cmd.todoItem)
				cmd.replyChan <- s.todos.Items
			case DeleteCommand:
				fmt.Println("DeleteCommand")
				s.todos.DeleteTodoItem(cmd.todoItem.Id)
				cmd.replyChan <- s.todos.Items
			default:
				log.Fatal("unknown command type", cmd.ct)
			}
		}
	}()
	return cmds
}

type Server struct {
	cmds chan<- Command
}

func (s *Server) todosHandler(w http.ResponseWriter, r *http.Request) {
	replyChan := make(chan []todos.Todo)
	switch r.Method {
	case http.MethodPost:
		log.Printf("POST %v", r)
		var todo todos.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
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
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
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
		log.Printf("GET %v", r)
		s.cmds <- Command{ct: DeleteCommand, todoItem: todos.Todo{Id: id}, replyChan: replyChan}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	reply := <-replyChan

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)

}
