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
	var mu sync.Mutex
	spySleeper := &SpySleeper{}
	buffer := &bytes.Buffer{}
	valuesChannel := make(chan int)

	go addNumbers(buffer, valuesChannel, &value, &mu)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			queueNumber(&wg, spySleeper, true, valuesChannel)
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			queueNumber(&wg, spySleeper, false, valuesChannel)
		}()
	}

	wg.Wait()

	// The final length should be 2000 if there are no race conditions
	if len(value) != 2000 {
		t.Errorf("expected value length to be 2000, but got %d", len(value))
	}
}
