package main

import (
	"bytes"
	"slices"
	"testing"
)

// write cities to file

// read cities from file

func TestSortStrings(t *testing.T) {
	got := sortStrings([]string{"b", "d", "c", "a"})
	want := []string{"a", "b", "c", "d"}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDisplaySlice(t *testing.T) {
	buffer := &bytes.Buffer{}
	displaySlice(buffer, []string{"hello", "world"})
	got := buffer.String()
	want := `hello
world
`

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
