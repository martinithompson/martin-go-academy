package main

import (
	"fmt"
	"testing"
)

func TestJoinStrings(t *testing.T) {
	got := JoinStrings([]string{"Hi", "how", "are", "you?"})
	want := "Hi how are you?"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleJoinStrings() {
	joined := JoinStrings([]string{"Hello world!"})
	fmt.Println(joined)
	// Output: Hello world!
}
