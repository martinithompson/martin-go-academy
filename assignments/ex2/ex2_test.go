package main

import (
	"fmt"
	"testing"
)

func TestDisplayName(t *testing.T) {
	got := DisplayName("Martin", "Ian", "Thompson")
	want := "Martin Ian Thompson"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleDisplayName() {
	fullName := DisplayName("Martin", "Ian", "Thompson")
	fmt.Println(fullName)
	// Output: Martin Ian Thompson
}
