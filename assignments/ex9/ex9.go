package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func getJoinedName(f, m, l string) string {
	return f + " " + m + " " + l
}

func splitText(text string) []string {
	return strings.Fields(text)
}

func DisplaySplitName(out io.Writer, name []string) {
	fmt.Fprintf(out, "first-name: %s\n", name[0])
	fmt.Fprintf(out, "middle-name: %s\n", name[1])
	fmt.Fprintf(out, "surname: %s\n", name[2])
}

func getName(in *bufio.Reader, out io.Writer, nameType string) string {
	fmt.Fprintf(out, "Please enter your %v name: ", nameType)
	name, _ := in.ReadString('\n')
	return strings.Replace(name, "\n", "", -1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstName := getName(reader, os.Stdout, "first")
	middleName := getName(reader, os.Stdout, "middle")
	lastName := getName(reader, os.Stdout, "last")

	fullName := getJoinedName(firstName, middleName, lastName)

	DisplaySplitName(os.Stdout, splitText(fullName))
}
