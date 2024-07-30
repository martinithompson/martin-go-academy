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

type Generator struct {
	Value []int
}

func (g *Generator) AddNumber(out io.Writer, s Sleeper, even bool) {
	s.Sleep()
	var rn int
	if even {
		rn = rand.Intn(51) * 2
	} else {
		rn = rand.Intn(50)*2 + 1
	}
	g.Value = append(g.Value, rn)
	fmt.Fprintf(out, "Updated value: %v\n", g.Value)
}

func main() {
	generator := Generator{}
	sleeper := DefaultSleeper{}
	fmt.Println("Generating...")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			generator.AddNumber(os.Stdout, &sleeper, true)
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			generator.AddNumber(os.Stdout, &sleeper, false)
		}()
	}

	wg.Wait()
	fmt.Printf("Final value: %v\n", generator.Value)
	fmt.Println("Done.")
}
