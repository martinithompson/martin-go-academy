package main

import (
	"fmt"
	"time"

	"github.com/bearbin/go-age"
)

type CalendarPeriod string

const (
	Day   CalendarPeriod = "day"
	Month CalendarPeriod = "month"
	Year  CalendarPeriod = "year"
)

func getCalendarPeriod(cp CalendarPeriod) int {
	var periodValue int
	fmt.Printf("Please enter the %v of your birth: ", cp)
	fmt.Scanf("%d", &periodValue)

	return periodValue
}

func getAge(birthday time.Time) int {
	return age.AgeAt(birthday, time.Now())
}

func getBirthdayAsDate(day, month, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func main() {
	day := getCalendarPeriod(Day)
	month := getCalendarPeriod(Month)
	year := getCalendarPeriod(Year)

	if day != 0 && month != 0 && year != 0 {
		fmt.Println("Your age is", getAge(getBirthdayAsDate(day, month, year)))
	}
}
