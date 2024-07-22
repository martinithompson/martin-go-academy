package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DisplayName(f, m, l string) string {
	return f + " " + m + " " + l
}

func getName(nameType string, r *bufio.Reader) string {
	fmt.Printf("Please enter your %v name: ", nameType)
	name, _ := r.ReadString('\n')
	return strings.Replace(name, "\n", "", -1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstName := getName("first", reader)
	middleName := getName("middle", reader)
	lastName := getName("last", reader)

	fmt.Printf("Your full name is %v\n", DisplayName(firstName, middleName, lastName))
}
