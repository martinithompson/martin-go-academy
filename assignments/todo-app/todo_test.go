package main

import (
	"bytes"
	"testing"
)

var washCar = Todo{Item: "Wash the car"}
var bookHol = Todo{Item: "Book holiday"}
var todos = Todos{items: []Todo{washCar, bookHol}}

func TestTodo(t *testing.T) {
	t.Run("todo description", func(t *testing.T) {
		got := washCar.Description()
		want := "Item: \"Wash the car\", Completed: no"
		assertStrings(t, got, want)
	})
}

func TestTodos(t *testing.T) {
	t.Run("todos print descriptions", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos.PrintDescriptions(buffer)
		got := buffer.String()
		want := "Item: \"Wash the car\", Completed: no\nItem: \"Book holiday\", Completed: no\n"
		assertStrings(t, got, want)
	})
	t.Run("todos output json", func(t *testing.T) {
		got := todos.OutputJson()
		want := "[{\"Item\":\"Wash the car\",\"Completed\":false},{\"Item\":\"Book holiday\",\"Completed\":false}]"

		assertStrings(t, got, want)
	})
}

func TestFormatCompleted(t *testing.T) {
	t.Run("returns 'yes' for true", func(t *testing.T) {
		got := formatCompleted(true)
		want := "yes"
		assertStrings(t, got, want)
	})
	t.Run("returns 'no' for false", func(t *testing.T) {
		got := formatCompleted(false)
		want := "no"
		assertStrings(t, got, want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
