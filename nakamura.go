/*
Package nakamura is a lightweight golang date library for parsing, validating, manipulating, and formatting dates
supporting both slash style and hyphen style dates
go get -u github.com/geniusmartinss/nakamura
nakamura.NewDate("2018-02-12", "YYYY-MM-DD") //{2018-02-12, "YYYY-MM-DD"}
nakamura.NewDate("", "YYYY-MM-DD") //Returns the date for the current day
*/

package nakamura

import (
	"fmt"
	"strings"
)

// Unit represents the date component to add or subtract.
type Unit int

const (
	// Year is the year component.
	Year Unit = iota
	// Month is the month component.
	Month
	// Day is the day component.
	Day
)

type Nakamura struct {
	date, format string
}

// NewDate creates a new Nakamura from a date string and format string.
// If date is empty or whitespace-only, it returns today's date in YYYY-MM-DD format.
// Returns an error if date is a datetime string (contains a space or 'T' character).
func NewDate(date, format string) (Nakamura, error) {
	if len(strings.TrimSpace(date)) == 0 {
		return Nakamura{Today(), "YYYY-MM-DD"}, nil
	}
	if strings.ContainsAny(date, " T") {
		return Nakamura{}, fmt.Errorf("nakamura: NewDate does not accept datetime strings, got %q", date)
	}
	return Nakamura{date, format}, nil
}

// IsDateValid checks for the validity of a nakamura date object
func (date Nakamura) IsDateValid() bool {
	return IsDateValid(date.date, date.format)
}

// Normalise validates the date and, if valid, returns a new Nakamura holding
// its canonical YYYY-MM-DD representation. It returns an error if the date is
// invalid (e.g. month 13, day 32, or February 30).
func (date Nakamura) Normalise() (Nakamura, error) {
	result, err := Normalise(date.date, date.format)
	if err != nil {
		return Nakamura{}, err
	}
	return Nakamura{result, date.format}, nil
}

// Humanise converts a nakamura date object into readable format
func (date Nakamura) Humanise() string {
	return Humanise(date.date, date.format)
}

// IsWeekEnd Checks if a given date input is a weekend
func (date Nakamura) IsWeekEnd() bool {
	return IsWeekend(date.date, date.format)
}

// IsLeapYear checks if the year in a nakamura date object is a leap year
func (date Nakamura) IsLeapYear() bool {
	return IsLeapYear(date.date, date.format)
}

// GreaterThan checks if firstDate is greater than secondDate
func (firstDate Nakamura) GreaterThan(secondDate Nakamura) bool {
	return GreaterThan(firstDate, secondDate, secondDate.format)
}

// LessThan checks if firstDate is Less than secondDate
func (firstDate Nakamura) LessThan(secondDate Nakamura) bool {
	return LessThan(firstDate, secondDate, secondDate.format)
}

// Check if firstDate is between secondDate and thirdDate
func (firstDate Nakamura) Between(secondDate, thirdDate Nakamura) bool {
	return firstDate.GreaterThan(secondDate) && firstDate.LessThan(thirdDate) || firstDate.Equal(secondDate) || firstDate.Equal(thirdDate)
}

// Check if a nakamura object is in the future
func (date Nakamura) IsFuture() bool {
	today, _ := NewDate("", date.format)
	return date.GreaterThan(today)
}

// Check if a nakamura object is in the past
func (date Nakamura) IsPast() bool {
	today, _ := NewDate("", date.format)
	return date.LessThan(today)
}

// Add adds a given value to the specified Unit component (Year, Month, or Day) of the date.
// Add returns an error for an unrecognised Unit value.
func (date Nakamura) Add(value int, unit Unit) (Nakamura, error) {
	return Add(date, value, unit)
}

// Subtract subtracts a given value from the specified Unit component (Year, Month, or Day) of the date.
// Subtract returns an error for an unrecognised Unit value.
func (date Nakamura) Subtract(value int, unit Unit) (Nakamura, error) {
	return Add(date, -value, unit)
}

// Equal checks if a given pair of date objects are equal
func (firstDate Nakamura) Equal(secondDate Nakamura) bool {
	return Equal(firstDate, secondDate, secondDate.format)
}

// Weekday checks if a given date falls on a weekday
func (date Nakamura) Weekday() string {
	return Weekday(date.date, date.format)
}

// Month returns the Month of a given date
func (date Nakamura) Month() string {
	return monthName(date.date, date.format)
}

func (date Nakamura) MonthDays() (int, error) {
	return DaysInMonth(date.date, date.format)
}

// Max returns the latest (maximum) date from the provided dates.
// If no dates are provided, it returns today's date in YYYY-MM-DD format.
func Max(dates ...Nakamura) Nakamura {
	if len(dates) > 0 {
		result := dates[0]
		for i := 1; i < len(dates); i++ {
			if GreaterThan(dates[i], result, result.format) {
				result = dates[i]
			}
		}
		return result
	}
	return Nakamura{Today(), "YYYY-MM-DD"}
}

// Min returns the earliest (minimum) date from the provided dates.
// If no dates are provided, it returns today's date in YYYY-MM-DD format.
func Min(dates ...Nakamura) Nakamura {
	if len(dates) > 0 {
		result := dates[0]
		for i := 1; i < len(dates); i++ {
			if LessThan(dates[i], result, result.format) {
				result = dates[i]
			}
		}
		return result
	}
	return Nakamura{Today(), "YYYY-MM-DD"}
}
