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

func TestGetRollResult(t *testing.T) {
	assertResult :=
		func(t testing.TB, got, want string) {
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		}
	t.Run("snake eyes", func(t *testing.T) {
		got := getRollResult(2)
		want := "Roll 2: SNAKE-EYES-CRAPS"
		assertResult(t, got, want)
	})
	t.Run("neutral", func(t *testing.T) {
		got := getRollResult(4)
		want := "Roll 4: NEUTRAL"
		assertResult(t, got, want)
	})
}
