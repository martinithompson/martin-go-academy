package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	todos "todo-app/project/todos"
)

func TestTodosHandler(t *testing.T) {
	server := Server{}
	server.cmds = startTodoManager()
	t.Run("GET todos when empty", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/todos", nil)

		rec := httptest.NewRecorder()
		server.todosHandler(rec, req)

		assertStatusCode(t, rec.Code, http.StatusOK)
		assertValue(t, rec.Body.String(), "[]\n")
	})
	t.Run("POST a todo", func(t *testing.T) {
		postName := "Test post"
		req := addTodo(postName)

		rec := httptest.NewRecorder()
		server.todosHandler(rec, req)

		var res []todos.Todo
		json.Unmarshal(rec.Body.Bytes(), &res)

		assertStatusCode(t, rec.Code, http.StatusCreated)
		assertValue(t, res[0].Name, postName)
		assertValue(t, res[0].Completed, false)
		assertValue(t, len(res), 1)
	})
	// t.Run("PATCH a todo", func(t *testing.T) {
	// })
	// t.Run("DELETE a todo", func(t *testing.T) {
	// })
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("incorrect status code: got %d want %d", got, want)
	}
}

func assertValue[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func addTodo(name string) *http.Request {
	n := map[string]string{"name": name}
	j, _ := json.Marshal(n)
	req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(j))
	return req
}

// func patchTodo(id, name, completed string) *http.Request {
// 	t := map[string]string{"id": id, "name": name, "completed": completed}
// 	j, _ := json.Marshal(t)
// 	req, _ := http.NewRequest(http.MethodPatch, "/todos/"+id, bytes.NewBuffer(j))
// 	return req
// }
