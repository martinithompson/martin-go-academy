package main

import (
	"encoding/json"
	"net/http"
	todos "todo-app/project/todos"
)

const jsonContentType = "application/json"

type TodoStore interface {
	AddTodo(todo todos.Todo)
	GetTodos() []todos.Todo
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

	t.Handler = router

	return t
}

func (t *TodoServer) postTodo(w http.ResponseWriter, todo todos.Todo) {
	t.store.AddTodo(todo)
	w.WriteHeader(http.StatusCreated)
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
		t.postTodo(w, todo)
	case http.MethodGet:
		t.getTodos(w)
	}
}
