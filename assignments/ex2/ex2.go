package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func DisplayName(out io.Writer, f, m, l string) {
	fmt.Fprintf(out, "Your full name is "+f+" "+m+" "+l)
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

	DisplayName(os.Stdout, firstName, middleName, lastName)
}
