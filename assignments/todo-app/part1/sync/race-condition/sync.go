package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(2 * time.Second)
}

func generateNumber(s Sleeper, even bool) int {
	s.Sleep()
	var rn int
	if even {
		rn = rand.Intn(51) * 2
	} else {
		rn = rand.Intn(50)*2 + 1
	}
	return rn
}

func addNumber(out io.Writer, wg *sync.WaitGroup, sleeper Sleeper, even bool, value *[]int) {
	defer wg.Done()
	gn := generateNumber(sleeper, even)
	*value = append(*value, gn)
	fmt.Fprintf(out, "Final value: %v\n", *value)
}

func main() {
	value := []int{}
	sleeper := DefaultSleeper{}
	var wg sync.WaitGroup

	fmt.Println("Generating...")

	// add even numbers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go addNumber(os.Stdout, &wg, &sleeper, true, &value)
	}

	// add odd numbers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go addNumber(os.Stdout, &wg, &sleeper, false, &value)
	}

	wg.Wait()
	fmt.Printf("Final value: %v\n", value)
	fmt.Println("Done.")
}
