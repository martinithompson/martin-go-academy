package main

import "testing"

func TestDisplayTodo(t *testing.T) {
	todo := Todo{Item: "Wash the car"}
	got := todo.Display()
	want := "Item: \"Wash the car\", Completed: no"
	assertStrings(t, got, want)
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
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
