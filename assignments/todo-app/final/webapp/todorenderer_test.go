package main

import (
	"bytes"
	"io"
	"testing"
	"todo-app/project/todos"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var washCar = todos.Todo{Name: "Wash the car", Id: "1"}
	var bookHol = todos.Todo{Name: "Book holiday", Id: "2"}
	var todoList = todos.TodoList{Items: []todos.Todo{washCar, bookHol}}

	todoRenderer, err := NewTodoRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it renders a list of todos", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := todoRenderer.List(&buf, todoList.Items); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
	t.Run("it renders a friendly message when there are no todos", func(t *testing.T) {
		noTodos := todos.TodoList{}
		buf := bytes.Buffer{}
		if err := todoRenderer.List(&buf, noTodos.Items); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
	t.Run("it renders an add page for adding a new todo", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := todoRenderer.Add(&buf); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
	t.Run("it renders an edit page for editing an existing todo", func(t *testing.T) {
		buf := bytes.Buffer{}
		editTodo := todos.Todo{Name: "Wash the car", Id: "1"}
		if err := todoRenderer.Edit(&buf, editTodo); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
	t.Run("it renders a delete page for deleting an existing todo", func(t *testing.T) {
		buf := bytes.Buffer{}
		deleteTodo := todos.Todo{Name: "Wash the car", Id: "1"}
		if err := todoRenderer.Delete(&buf, deleteTodo); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var washCar = todos.Todo{Name: "Wash the car", Id: "1"}
	var bookHol = todos.Todo{Name: "Book holiday", Id: "2"}
	var todoList = todos.TodoList{Items: []todos.Todo{washCar, bookHol}}

	todoRenderer, _ := NewTodoRenderer()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		todoRenderer.List(io.Discard, todoList.Items)
	}
}
