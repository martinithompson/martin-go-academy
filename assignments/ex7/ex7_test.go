package main

import (
	"bytes"
	"testing"
)

type SpyRandomiser struct{}

func (s SpyRandomiser) Generate(_ int) int {
	return 4
}

func TestRollDice(t *testing.T) {
	randomiser := SpyRandomiser{}
	got := rollDice(randomiser)
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestRunRolls(t *testing.T) {
	randomiser := SpyRandomiser{}
	buffer := &bytes.Buffer{}
	runRolls(1, randomiser, buffer)

	want := "Roll 8: NEUTRAL\n"

	if buffer.String() != want {
		t.Errorf("got %q want %q", buffer.String(), want)
	}
}

func TestGetRollResult(t *testing.T) {
	assertResult :=
		func(t testing.TB, got, want string) {
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		}
	t.Run("snake eyes", func(t *testing.T) {
		got := getRollResult(2)
		want := "Roll 2: SNAKE-EYES-CRAPS\n"
		assertResult(t, got, want)
	})
	t.Run("neutral", func(t *testing.T) {
		got := getRollResult(4)
		want := "Roll 4: NEUTRAL\n"
		assertResult(t, got, want)
	})
}
