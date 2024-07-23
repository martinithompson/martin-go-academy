package main

import (
	"slices"
	"testing"
)

func TestPopulateArray(t *testing.T) {
	got := populateArray()
	want := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIsEven(t *testing.T) {
	checkValues := func(t testing.TB, got, want bool) {
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("returns true for an even number", func(t *testing.T) {
		got := isEven(2)
		want := true

		checkValues(t, got, want)
	})
	t.Run("returns false for an odd number", func(t *testing.T) {
		got := isEven(7)
		want := false

		checkValues(t, got, want)
	})
}

func TestAppendEvenOrOdd(t *testing.T) {
	checkAppendedSlices := func(t testing.TB, got, want []int, parity string) {
		if !slices.Equal(got, want) {
			t.Errorf("for %v value got slice %v want slice %v", parity, got, want)
		}
	}
	t.Run("appends an even value as expected", func(t *testing.T) {
		gotEvens, gotOdds := appendEvenOrOdd([]int{2}, []int{3}, 4)
		wantEvens := []int{2, 4}
		wantOdds := []int{3}

		checkAppendedSlices(t, gotEvens, wantEvens, "even")
		checkAppendedSlices(t, gotOdds, wantOdds, "odd")
	})
	t.Run("appends an odd value as expected", func(t *testing.T) {
		gotEvens, gotOdds := appendEvenOrOdd([]int{2}, []int{3}, 7)
		wantEvens := []int{2}
		wantOdds := []int{3, 7}

		checkAppendedSlices(t, gotEvens, wantEvens, "even")
		checkAppendedSlices(t, gotOdds, wantOdds, "odd")
	})
}
