package main

import "testing"

func TestNumberOfDigits(t *testing.T) {
	t.Run("single digit", func(t *testing.T) {
		got := numberOfDigits(1)
		want := 1
		assertDigits(t, got, want)
	})
	t.Run("two digits", func(t *testing.T) {
		got := numberOfDigits(16)
		want := 2
		assertDigits(t, got, want)
	})
	t.Run("three digits", func(t *testing.T) {
		got := numberOfDigits(641)
		want := 3
		assertDigits(t, got, want)
	})
}

func TestSumArray(t *testing.T) {
	got := sumArray([3]int{1, 2, 3})
	want := 6
	assertDigits(t, got, want)
}

func assertDigits(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
