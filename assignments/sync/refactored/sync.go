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

func queueNumber(wg *sync.WaitGroup, sleeper Sleeper, even bool, ch chan<- int) {
	defer wg.Done()
	gn := generateNumber(sleeper, even)
	ch <- gn
}

func addNumbers(out io.Writer, ch <-chan int, value *[]int, mu *sync.Mutex) {
	for num := range ch {
		mu.Lock()
		*value = append(*value, num)
		fmt.Fprintf(out, "Current value: %v\n", *value)
		mu.Unlock()
	}
}

func main() {
	value := []int{}
	sleeper := DefaultSleeper{}
	var wg sync.WaitGroup
	var mu sync.Mutex
	valuesChannel := make(chan int)

	fmt.Println("Generating...")

	go addNumbers(os.Stdout, valuesChannel, &value, &mu)

	// Add even numbers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go queueNumber(&wg, &sleeper, true, valuesChannel)
	}

	// Add odd numbers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go queueNumber(&wg, &sleeper, false, valuesChannel)
	}

	wg.Wait()
	close(valuesChannel)

	fmt.Printf("Final value: %v\n", value)
	fmt.Println("Done.")
}
