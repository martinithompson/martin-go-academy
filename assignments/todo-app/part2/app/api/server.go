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
	t := new(TodoServer)

	t.store = store

	router := http.NewServeMux()
	router.Handle("/todos", http.HandlerFunc(t.todosHandler))
	router.Handle("/todos/", http.HandlerFunc(t.todoHandler))

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
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		log.Printf("Create todo: %v", todo)
		t.postTodo(w, todo)
	case http.MethodGet:
		log.Println("Get todos")
		t.getTodos(w)
	}
}

func (t *TodoServer) todoHandler(w http.ResponseWriter, r *http.Request) {
	todoIndex := strings.TrimPrefix(r.URL.Path, "/todos/")

	todoInt, _ := strconv.Atoi(todoIndex)
	todoInt++

	switch r.Method {
	case http.MethodPatch:
		var todo todos.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		log.Printf("Edit todo %d: %v", todoInt, todo)
		t.editTodo(w, todoInt, todo)
	case http.MethodDelete:
		log.Printf("Delete todo: %d", todoInt)
		t.deleteTodo(w, todoInt)
	}
}
