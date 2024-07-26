package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/bearbin/go-age"
)

type Pupil struct {
	Fullname    string
	DateOfBirth time.Time
}

func (p Pupil) FormattedDateOfBirth() string {
	year, month, day := p.DateOfBirth.Date()
	return fmt.Sprintf("%02d-%02d-%04d", day, int(month), year)
}

func (p Pupil) Age() int {
	return age.AgeAt(p.DateOfBirth, time.Now())
}

func DisplayPupil(out io.Writer, p Pupil) {
	fmt.Fprintf(out, "full-name: %v, dob: %v, age: %d\n", p.Fullname, p.FormattedDateOfBirth(), p.Age())
}

func ReadPupils() (pupils []string) {
	fi, err := os.Open("pupils.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(fi)

	var line []byte
	var readErr error

	for readErr == nil {
		line, _, readErr = r.ReadLine()
		pupils = append(pupils, string(line))
	}
	return
}

func convertDateStringToTime(inputDate string) (time.Time, error) {
	layout := "02-01-2006"
	date, err := time.Parse(layout, inputDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func main() {
	pupilsData := ReadPupils()
	var pupils []Pupil
	var pupilSlice []string
	var birthDate time.Time
	var name string

	for _, pd := range pupilsData {
		pupilSlice = strings.Split(pd, ",")
		if len(pupilSlice) == 2 {
			name = pupilSlice[0]
			birthDate, _ = convertDateStringToTime(pupilSlice[1])

			pupils = append(pupils, Pupil{name, birthDate})
		}

	}

	for _, p := range pupils {
		DisplayPupil(os.Stdout, p)
	}

}
