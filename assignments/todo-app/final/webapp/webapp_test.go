package main

import (
	"bytes"
	"testing"
	"todo-app/project/todos"
)

func TestRenderList(t *testing.T) {

	var (
		aTodo = todorenderer.Todo{
			Name:      "hello world",
			Completed: false,
		}
	)
	t.Run("it renders no posts when the todo list is empty", func(t *testing.T) {
		buf := bytes.Buffer{}
		emptyTodos := todos.TodoList{}
		err := todorenderer.Render(&buf, emptyTodos)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world<h1>`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
