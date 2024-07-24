package main

import (
	"bytes"
	"os"
	"testing"
)

func TestDisplayName(t *testing.T) {
	buffer := &bytes.Buffer{}
	DisplayName(buffer, "Martin", "Ian", "Thompson")
	want := "Your full name is Martin Ian Thompson"

	if buffer.String() != want {
		t.Errorf("got %q want %q", buffer.String(), want)
	}
}

func ExampleDisplayName() {
	DisplayName(os.Stdout, "Martin", "Ian", "Thompson")
	// Output: Your full name is Martin Ian Thompson
}
