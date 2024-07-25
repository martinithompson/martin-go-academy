package main

import (
	"bytes"
	"slices"
	"testing"
)

func TestSplitText(t *testing.T) {
	got := splitText("Bob The Builder")
	want := []string{"Bob", "The", "Builder"}

	if !slices.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDisplaySplitName(t *testing.T) {
	buffer := &bytes.Buffer{}
	DisplaySplitName(buffer, []string{"Bob", "The", "Builder"})
	got := buffer.String()
	want := `first-name: Bob
middle-name: The
surname: Builder
`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetJoinedName(t *testing.T) {
	got := getJoinedName("Bob", "The", "Builder")
	want := "Bob The Builder"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
