package main

import "testing"

type SpyRandomiser struct{}

func (s SpyRandomiser) Generate(_ int) int {
	return 4
}

func TestRollDice(t *testing.T) {
	randomiser := SpyRandomiser{}
	got := RollDice(randomiser)
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// test create array of 2 dice rolls

// test roll x(50) times - store in array of scores

// test loop through array and output score, NATURAL etc
