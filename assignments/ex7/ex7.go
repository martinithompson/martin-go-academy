package main

import (
	"fmt"
	"math/rand"
)

type Randomiser interface {
	Generate(int) int
}

type IntRandomiser struct{}

func (r IntRandomiser) Generate(max int) int {
	return rand.Intn(max) + 1
}

func RollDice(r Randomiser) int {
	return r.Generate(6)
}

func main() {
	randomiser := IntRandomiser{}
	fmt.Println(RollDice(randomiser))
}
