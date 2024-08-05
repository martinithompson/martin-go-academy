package todos

import (
	"bytes"
	"slices"
	"testing"
	"testing/fstest"
)

var washCar = Todo{Name: "Wash the car", Id: "1"}
var bookHol = Todo{Name: "Book holiday", Id: "2"}
var todos = TodoList{Items: []Todo{washCar, bookHol}}

func TestTodo(t *testing.T) {
	t.Run("todo description", func(t *testing.T) {
		got := washCar.Description()
		want := "Name: \"Wash the car\", Completed: no"
		assertStrings(t, got, want)
	})
}

func TestTodos(t *testing.T) {
	t.Run("todos print descriptions", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos.PrintDescriptions(buffer)
		got := buffer.String()
		want := "\t1> Name: \"Wash the car\", Completed: no\n\t2> Name: \"Book holiday\", Completed: no\n"
		assertStrings(t, got, want)
	})
	t.Run("todos output json", func(t *testing.T) {
		got := todos.Json()
		want := "[{\"id\":\"1\",\"name\":\"Wash the car\",\"completed\":false},{\"id\":\"2\",\"name\":\"Book holiday\",\"completed\":false}]"

		assertStrings(t, got, want)
	})
	t.Run("todos save", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos.Save(buffer)

		want := "[{\"id\":\"1\",\"name\":\"Wash the car\",\"completed\":false},{\"id\":\"2\",\"name\":\"Book holiday\",\"completed\":false}]"
		assertStrings(t, buffer.String(), want)
	})
	t.Run("todos add todo", func(t *testing.T) {
		gotTodos := TodoList{}
		gotTodos.AddTodos(washCar, bookHol)

		assertTodos(t, gotTodos, todos)
	})
	t.Run("todos add todo as task string", func(t *testing.T) {
		gotTodos := TodoList{}
		gotTodos.AddTodosByName("Wash the car", "Book holiday")

		assertStrings(t, gotTodos.Items[0].Name, todos.Items[0].Name)
		assertStrings(t, gotTodos.Items[1].Name, todos.Items[1].Name)
	})
	t.Run("todos delete", func(t *testing.T) {
		gotTodos := TodoList{Items: []Todo{washCar}}
		gotTodos.DeleteTodo("1")

		assertTodos(t, gotTodos, TodoList{})
	})
	t.Run("todos edit", func(t *testing.T) {
		gotTodos := TodoList{Items: []Todo{washCar}}
		gotTodos.UpdateTodo(Todo{Name: "Paint the fence", Id: "1", Completed: true})

		assertStrings(t, gotTodos.Items[0].Name, "Paint the fence")
		assertCompleted(t, gotTodos.Items[0].Completed, true)
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
	json := "[{\"id\":\"1\",\"name\":\"Wash the car\",\"completed\":false},{\"id\":\"2\",\"name\":\"Book holiday\",\"completed\":false}]"
	fs := fstest.MapFS{
		"todos.json": {Data: []byte(json)},
	}
	newTodos := TodoList{}
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

func assertCompleted(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func assertTodos(t *testing.T, got, want TodoList) {
	t.Helper()
	if !slices.Equal(got.Items, want.Items) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
