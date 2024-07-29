package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func NumberGenerator(out io.Writer, sleeper Sleeper, wg *sync.WaitGroup, amount int, values *[]int, evenNumbers bool) {
	defer wg.Done()
	start := 1
	if evenNumbers {
		start = 2
	}
	for i := start; i <= amount*2; i += 2 {
		fmt.Fprintf(out, "Generated: %d\n", i)
		sleeper.Sleep()
		*values = append(*values, i)
	}
}

func main() {
	var wg sync.WaitGroup
	var values []int
	sleeper := DefaultSleeper{}

	wg.Add(2)
	go NumberGenerator(os.Stdout, &sleeper, &wg, 5, &values, true)
	go NumberGenerator(os.Stdout, &sleeper, &wg, 5, &values, false)
	fmt.Println("Running...")
	wg.Wait()
	fmt.Printf("Output: %v\n", values)
}
