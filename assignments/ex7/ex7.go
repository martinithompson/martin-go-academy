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

func SumTwoDice() int {
	randomiser := IntRandomiser{}
	diceOne := RollDice(randomiser)
	diceTwo := RollDice(randomiser)
	return diceOne + diceTwo
}

func getRollResult(total int) string {
	var result string
	switch total {
	case 7, 11:
		result = fmt.Sprintf("Roll %d: NATURAL", total)
	case 2:
		result = fmt.Sprintf("Roll %d: SNAKE-EYES-CRAPS", total)
	case 3, 12:
		result = fmt.Sprintf("Roll %d: LOSS-CRAPS", total)
	default:
		result = fmt.Sprintf("Roll %d: NEUTRAL", total)
	}
	return result
}

func RollMultiple(count int) {
	var total int
	for i := 1; i <= count; i++ {
		total = SumTwoDice()
		fmt.Println(getRollResult(total))
	}
}

func main() {
	RollMultiple(50)
}
