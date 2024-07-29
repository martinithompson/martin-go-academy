package main

import (
	"bytes"
	"slices"
	"sync"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestNumberGenerator(t *testing.T) {
	assertEqual := func(t *testing.T, got, want []int) {
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	runTest := func(t *testing.T, amount int, even bool, want []int) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		wg := &sync.WaitGroup{}
		wg.Add(1)
		value := []int{}
		NumberGenerator(buffer, spySleeper, wg, amount, &value, even)
		wg.Wait()
		assertEqual(t, value, want)

		if spySleeper.Calls != amount {
			t.Errorf("not enough calls to sleeper, want %d got %d", amount, spySleeper.Calls)
		}
	}
	t.Run("even numbers", func(t *testing.T) {
		runTest(t, 5, true, []int{2, 4, 6, 8, 10})
	})
	t.Run("odd numbers", func(t *testing.T) {
		runTest(t, 5, false, []int{1, 3, 5, 7, 9})
	})
}
