package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	todos "todo-app/project/todos"
)

const jsonContentType = "application/json"

type TodoStore interface {
	AddTodo(todo todos.Todo)
	GetTodos() []todos.Todo
	DeleteTodo(index int)
	EditTodo(index int, todo todos.Todo)
}

type TodoServer struct {
	store TodoStore
	http.Handler
}

func NewTodoServer(store TodoStore) *TodoServer {
	t := &TodoServer{store: store}
	router := http.NewServeMux()
	router.HandleFunc("/todos", t.todosHandler)
	router.HandleFunc("/todos/", t.todoHandler)

	t.Handler = router
	return t
}

func (t *TodoServer) postTodo(w http.ResponseWriter, todo todos.Todo) {
	t.store.AddTodo(todo)
	w.WriteHeader(http.StatusCreated)
}

func (t *TodoServer) editTodo(w http.ResponseWriter, index int, todo todos.Todo) {
	t.store.EditTodo(index, todo)
	w.WriteHeader(http.StatusCreated)
}

func (t *TodoServer) deleteTodo(w http.ResponseWriter, index int) {
	t.store.DeleteTodo(index)
	w.WriteHeader(http.StatusNoContent)
}

func (t *TodoServer) getTodos(w http.ResponseWriter) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(t.store.GetTodos())
}

func (t *TodoServer) todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var todo todos.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		log.Printf("Create todo: %v", todo)
		t.postTodo(w, todo)
	case http.MethodGet:
		log.Println("Get todos")
		t.getTodos(w)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (t *TodoServer) todoHandler(w http.ResponseWriter, r *http.Request) {
	todoIndex, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/todos/"))
	if err != nil {
		http.Error(w, "Invalid todo index", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPatch:
		var todo todos.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		log.Printf("Edit todo %d: %v", todoIndex, todo)
		t.editTodo(w, todoIndex, todo)
	case http.MethodDelete:
		log.Printf("Delete todo: %d", todoIndex)
		t.deleteTodo(w, todoIndex)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
