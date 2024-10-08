package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestMenu(t *testing.T) {
	t.Run("open menu", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		menu(buffer)
		want := `
	*** To-Do Options ***
	Please enter an option number to continue:
	
	1) Add a new to-do item
	2) View all to-dos
	3) Update a to-do item
	4) Delete a to-do item
	5) Exit
`
		assertStrings(t, buffer.String(), want)
	})
}
func TestReadOption(t *testing.T) {
	assertOption := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Fatalf("expected option to be %d, got %d", got, want)
		}
	}
	assertNoError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}
	assertError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Fatalf("expected error but did not get one")
		}
	}
	t.Run("valid integer", func(t *testing.T) {
		input := "4\n"
		reader := bufio.NewReader(strings.NewReader(input))
		buffer := &bytes.Buffer{}
		option, err := readOption(reader, buffer, 6)

		assertNoError(t, err)
		assertOption(t, option, 4)
	})
	t.Run("invalid integer", func(t *testing.T) {
		input := "7\n"
		reader := bufio.NewReader(strings.NewReader(input))
		buffer := &bytes.Buffer{}
		option, err := readOption(reader, buffer, 4)

		assertError(t, err)
		assertOption(t, option, 0)
	})
	t.Run("invalid string", func(t *testing.T) {
		input := "abc\n"
		reader := bufio.NewReader(strings.NewReader(input))
		buffer := &bytes.Buffer{}
		option, err := readOption(reader, buffer, 4)

		assertError(t, err)
		assertOption(t, option, 0)
	})
}

func TestReadItem(t *testing.T) {
	input := "new item\n"
	reader := bufio.NewReader(strings.NewReader(input))
	option, _ := readItem(reader)

	assertStrings(t, option, "new item")
}

func TestAddTodo(t *testing.T) {
	input := "new item\n"
	reader := bufio.NewReader(strings.NewReader(input))
	option, _ := readItem(reader)

	assertStrings(t, option, "new item")
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
