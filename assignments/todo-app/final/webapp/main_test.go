package main

import (
	"testing"
	"todo-app/project/todos"
)

func TestGetIdFromPath(t *testing.T) {
	id := "12345"
	got := getIdFromPath("/edit/" + id)

	assertValue(t, got, id)
}
func TestConvertCheckboxValueToBool(t *testing.T) {
	t.Run("returns true for on", func(t *testing.T) {
		got := convertCheckboxValueToBool("on")
		assertValue(t, got, true)
	})
	t.Run("returns false for off", func(t *testing.T) {
		got := convertCheckboxValueToBool("off")
		assertValue(t, got, false)
	})
	t.Run("returns false for empty string", func(t *testing.T) {
		got := convertCheckboxValueToBool("")
		assertValue(t, got, false)
	})
}

func TestGetTodoById(t *testing.T) {
	washCar := todos.Todo{Name: "Wash the car", Id: "1"}
	bookHol := todos.Todo{Name: "Book holiday", Id: "2"}
	localStore.Items = []todos.Todo{
		washCar,
		bookHol,
	}
	got := getTodoById(bookHol.Id)

	assertValue(t, got, bookHol)
}

func assertValue[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
