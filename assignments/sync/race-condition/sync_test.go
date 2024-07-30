package main

import (
	"bytes"
	"sync"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestGenerator(t *testing.T) {
	value := []int{}
	var wg sync.WaitGroup
	spySleeper := &SpySleeper{}
	buffer := &bytes.Buffer{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			addNumber(buffer, &wg, spySleeper, true, &value)
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			addNumber(buffer, &wg, spySleeper, false, &value)
		}()
	}

	wg.Wait()

	// The final length should be 2000 if there are no race conditions
	if len(value) != 2000 {
		t.Errorf("expected value length to be 2000, but got %d", len(value))
	}
}
