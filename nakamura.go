/*
Package nakamura is a lightweight golang date library for parsing, validating, manipulating, and formatting dates
supporting both slash style and hyphen style dates
go get -u github.com/geniusmartinss/nakamura
nakamura.NewDate("2018-02-12", "YYYY-MM-DD") //{2018-02-12, "YYYY-MM-DD"}
nakamura.NewDate("", "YYYY-MM-DD") //Returns the date for the current day
 */

package nakamura

import "strings"

type Nakamura struct {
	date, format string
}

// NewDate creates a new nakamura object to perform
//various date formatting and manipulations n
func NewDate(date, format string) Nakamura {

	if len(strings.TrimSpace(date)) == 0 {
		return Nakamura{Today(), "YYYY-MM-DD"}
	}
	return Nakamura{strings.Split(date, " ")[0], format}
}

// IsDateValid checks for the validity of a nakamura date object
func (date Nakamura) IsDateValid() bool {
	return IsDateValid(date.date, date.format)
}

// Normalise corrects any errors in the date object, such as
// overflowing days or months by normalising
func (date Nakamura) Normalise() Nakamura {
	return Nakamura{Normalise(date.date, date.format), date.format}
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
func (firstDate Nakamura)Between(secondDate, thirdDate Nakamura)  bool{
	return firstDate.GreaterThan(secondDate) && firstDate.LessThan(thirdDate) || firstDate.Equal(secondDate) || firstDate.Equal(thirdDate)
}

// Check if a nakamura object is in the future
func (date Nakamura)IsFuture()  bool{
	return date.GreaterThan(NewDate("", date.format))
}

// Check if a nakamura object is in the past
func (date Nakamura)IsPast()  bool{
	return date.LessThan(NewDate("", date.format))
}

// Add adds a given value to either year/month/day
func (date Nakamura) Add(value int, format string) Nakamura {
	return Add(date, value, format)
}

// Subtract subtracts a given value from year/month/day
func (date Nakamura) Subtract(value int, format string) Nakamura {
	return Add(date, -value, format)
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
	return Month(date.date, date.format)
}
