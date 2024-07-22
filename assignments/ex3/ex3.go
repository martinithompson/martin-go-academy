package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Limit struct {
	Min int
	Max int
}

func IsBetweenLimits(l Limit, value int) bool {
	return value >= l.Min && value <= l.Max
}

func getValue(valueType string, r *bufio.Reader) int {
	fmt.Printf("Enter your %v: ", valueType)
	valueStr, _ := r.ReadString('\n')
	trimmedStr := strings.Replace(valueStr, "\n", "", -1)
	value, _ := strconv.Atoi(trimmedStr)

	return value
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	min := getValue("minimum", reader)
	max := getValue("maximum", reader)
	value := getValue("value", reader)

	limit := Limit{Min: min, Max: max}

	if IsBetweenLimits(limit, value) {
		fmt.Printf("YES! Value %v is between the limits %v and %v\n", value, limit.Min, limit.Max)
	} else {
		fmt.Printf("NO! Value %v is not between the limits %v and %v\n", value, limit.Min, limit.Max)
	}
}
