package main

import (
	"fmt"
	"strings"
)

func JoinStrings(inputStrings []string) string {
	return strings.Join(inputStrings, " ")
}

func main() {
	myStrings := []string{"Hello", "world!"}
	fmt.Println(JoinStrings(myStrings))
}
