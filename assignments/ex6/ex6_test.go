package main

import (
	"testing"
	"time"
)

func TestAge(t *testing.T) {
	birthday := time.Date(1978, time.Month(5), 10, 0, 0, 0, 0, time.UTC)
	got := getAge(birthday)
	want := 46

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestGetBirthdayAsDate(t *testing.T) {
	got := getBirthdayAsDate(10, 5, 1978)
	want := time.Date(1978, time.Month(5), 10, 0, 0, 0, 0, time.UTC)

	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}
