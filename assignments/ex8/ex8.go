package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var cities = [...]string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteCities() {
	f, err := os.Create("./cities.txt")
	check(err)

	defer f.Close()

	var line string
	for _, city := range cities {
		line = fmt.Sprintf("%s\n", city)
		f.WriteString(line)
	}

	f.Sync()
}

func ReadCities() []string {
	fi, err := os.Open("cities.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(fi)

	var line []byte
	var readErr error
	var cities []string
	for readErr == nil {
		line, _, readErr = r.ReadLine()
		cities = append(cities, string(line))
	}
	return cities
}

func sortStrings(source []string) []string {
	sort.Strings(source)
	return source
}

func displaySlice(out io.Writer, source []string) {
	for _, line := range source {
		fmt.Fprintln(out, line)
	}
}

func main() {
	WriteCities()
	cities := sortStrings(ReadCities())
	displaySlice(os.Stdout, cities)
}
