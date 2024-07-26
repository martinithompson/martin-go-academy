package main

import (
	"bytes"
	"reflect"
	"testing"
	"testing/fstest"
)

var washCar = Todo{Item: "Wash the car"}
var bookHol = Todo{Item: "Book holiday"}
var todos = Todos{Items: []Todo{washCar, bookHol}}

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
		got := todos.Json()
		want := "[{\"Item\":\"Wash the car\",\"Completed\":false},{\"Item\":\"Book holiday\",\"Completed\":false}]"

		assertStrings(t, got, want)
	})
	t.Run("todos save", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos.Save(buffer)

		want := "[{\"Item\":\"Wash the car\",\"Completed\":false},{\"Item\":\"Book holiday\",\"Completed\":false}]"
		assertStrings(t, buffer.String(), want)
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

func TestLoadTodos(t *testing.T) {
	json := "[{\"Item\":\"Wash the car\",\"Completed\":false},{\"Item\":\"Book holiday\",\"Completed\":false}]"
	fs := fstest.MapFS{
		"todos.json": {Data: []byte(json)},
	}
	newTodos := Todos{}
	newTodos.Load(fs, "todos.json")

	got := len(newTodos.Items)
	want := 2

	if got != want {
		t.Errorf("got %d todos, wanted %d todos", got, want)
	}

	assertTodos(t, newTodos, todos)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}

func assertTodos(t *testing.T, got, want Todos) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
