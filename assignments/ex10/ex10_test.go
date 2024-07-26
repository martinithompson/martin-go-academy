package main

import (
	"bytes"
	"testing"
	"time"
)

const dobString = "14-02-1990"

var dob, _ = convertDateStringToTime(dobString)
var pupil = Pupil{"Joe Bloggs", dob}

func TestFormatDateOfBirth(t *testing.T) {
	got := pupil.FormattedDateOfBirth()

	if got != dobString {
		t.Errorf("got %q want %q,", got, dobString)
	}
}

func TestAge(t *testing.T) {
	got := pupil.Age()
	want := 34

	if got != want {
		t.Errorf("got %d want %d,", got, want)
	}
}

func TestDisplayPupil(t *testing.T) {
	buffer := &bytes.Buffer{}
	DisplayPupil(buffer, pupil)
	got := buffer.String()
	want := "full-name: Joe Bloggs, dob: 14-02-1990, age: 34\n"

	if got != want {
		t.Errorf("got %q want %q,", got, want)
	}
}

func TestConvertDateStringToTime(t *testing.T) {
	got, _ := convertDateStringToTime("01-03-2005")
	want := time.Date(2005, time.Month(3), 1, 0, 0, 0, 0, time.UTC)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
