package main

import (
	"bytes"
	"os"
	"testing"
)

func TestDisplayText(t *testing.T) {
	buffer := &bytes.Buffer{}
	DisplayText(buffer, "Martin Ian Thompson")
	want := "Martin Ian Thompson"

	if buffer.String() != want {
		t.Errorf("got %q want %q", buffer.String(), want)
	}
}

func TestGetFullName(t *testing.T) {
	got := getFullName("Bob", "The", "Builder")
	want := "Bob The Builder"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleDisplayText() {
	DisplayText(os.Stdout, "Martin Ian Thompson")
	// Output: Martin Ian Thompson
}
