package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

type Randomiser interface {
	Generate(int) int
}

type IntRandomiser struct{}

func (r IntRandomiser) Generate(max int) int {
	return rand.Intn(max) + 1
}

func rollDice(r Randomiser) int {
	return r.Generate(6)
}

func getRollResult(total int) string {
	var result string
	switch total {
	case 7, 11:
		result = fmt.Sprintf("Roll %d: NATURAL\n", total)
	case 2:
		result = fmt.Sprintf("Roll %d: SNAKE-EYES-CRAPS\n", total)
	case 3, 12:
		result = fmt.Sprintf("Roll %d: LOSS-CRAPS\n", total)
	default:
		result = fmt.Sprintf("Roll %d: NEUTRAL\n", total)
	}
	return result
}

func runRolls(numRolls int, r Randomiser, out io.Writer) {
	var total int
	for i := 1; i <= numRolls; i++ {
		total = rollDice(r) + rollDice(r)
		fmt.Fprint(out, getRollResult(total))
	}
}

func main() {
	intRandomiser := IntRandomiser{}
	runRolls(50, intRandomiser, os.Stdout)
}
