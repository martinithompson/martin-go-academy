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
	buffer := &bytes.Buffer{}
	g := &Generator{}
	var wg sync.WaitGroup
	spySleeper := &SpySleeper{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			g.AddNumber(buffer, spySleeper, true)
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			g.AddNumber(buffer, spySleeper, false)
		}()
	}

	wg.Wait()

	// The final length should be 2000 if there are no race conditions
	if len(g.Value) != 2000 {
		t.Errorf("expected value length to be 0, but got %d", len(g.Value))
	}
}
