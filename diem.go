package nakamura

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// daysInMonth contains the number of days in each month for a non-leap year.
// Index 0 is unused; months 1–12 map directly to daysInMonth[month].
var daysInMonth = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func IsDateValid(input, format string) bool {
	if len(input) == 0 {
		return false
	}
	date, dateFormat := getDateType(input, format)
	dateIntegerCount := 0

	if len(dateFormat) != 3 || len(date) != 3 {
		return false
	}

	for _, chunk := range date {
		if _, err := strconv.Atoi(chunk); err == nil {
			dateIntegerCount++
		}
	}

	if dateIntegerCount != 3 {
		return false
	}

	if !dateMatchesFormat(date, dateFormat) {
		return false
	}

	return true
}

//the split version of both
func dateMatchesFormat(input, format []string) bool {
	var globalMonth int
	var globalYear int
	for i := 0; i < len(input); i++ {
		if len(input[i]) != len(format[i]) {
			return false
		}
		if month, err := strconv.Atoi(input[i]); format[i] == "MM" && err == nil {
			globalMonth = month
			if month > 12 || month < 1 {
				return false
			}
		}
		if year, err := strconv.Atoi(input[i]); format[i] == "YYYY" && err == nil {
			globalYear = year
		}
		if day, err := strconv.Atoi(input[i]); format[i] == "DD" && err == nil {
			return isDayValidInMonth(globalYear, globalMonth, day)
		}

		//check if the day of a month is valid
	}
	return true
}

func isDayValidInMonth(year, month, day int) bool {
	if month < 1 || month > 12 {
		return false
	}
	days := daysInMonth[month]
	if month == 2 && isLeapYear(year) {
		days = 29
	}
	return day >= 1 && day <= days
}

// Normalise validates the input date and, if valid, returns its canonical
// YYYY-MM-DD string representation. It returns an error if the date is invalid
// (e.g. month 13, day 32, or February 30).
func Normalise(input, format string) (string, error) {
	date, dateFormat := getDateType(input, format)
	year, month, day := returnYearMonthDay(date, dateFormat)
	if !isDayValidInMonth(year, month, day) {
		return "", fmt.Errorf("nakamura: invalid date %q", input)
	}
	return strings.Split(
		time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).String(),
		" ")[0], nil
}

func returnYearMonthDay(input, format []string) (year, month, day int) {
	var myDay int
	var myMonth int
	var myYear int
	for i := 0; i < len(input); i++ {
		if year, err := strconv.Atoi(input[i]); format[i] == "YYYY" && err == nil {
			myYear = year
		}
		if month, err := strconv.Atoi(input[i]); format[i] == "MM" && err == nil {
			myMonth = month
		}
		if day, err := strconv.Atoi(input[i]); format[i] == "DD" && err == nil {
			myDay = day
		}
	}
	return myYear, myMonth, myDay
}

func Humanise(input, format string) string {
	date, dateFormat := getDateType(input, format)
	year, month, day := returnYearMonthDay(date, dateFormat)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Weekday().String() + "," + strconv.Itoa(day) + " " + getMonth(month).String() + " " + strconv.Itoa(year)
}

func getDateType(input, format string) (date, dateFormat []string) {
	//check if date is type hyphen or slash
	formatHyphen := strings.Split(format, "-")
	hyph := false
	if len(formatHyphen) == 3 {
		hyph = true
	}
	if hyph {
		date := strings.Split(input, "-")
		return date, formatHyphen
	} else {
		date := strings.Split(input, "/")
		formatSlash := strings.Split(format, "/")
		return date, formatSlash
	}

}

func getMonth(month int) time.Month {
	switch month {
	case 1:
		return time.January
	case 2:
		return time.February
	case 3:
		return time.March
	case 4:
		return time.April
	case 5:
		return time.May
	case 6:
		return time.June
	case 7:
		return time.July
	case 8:
		return time.August
	case 9:
		return time.September
	case 10:
		return time.October
	case 11:
		return time.November
	case 12:
		return time.December
	default:
		return time.January
	}
}

func Today() string {
	return strings.Split(time.Now().String(), " ")[0]
}

func IsWeekend(input, format string) bool {
	date, dateFormat := getDateType(input, format)
	year, month, day := returnYearMonthDay(date, dateFormat)

	if weekday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Weekday().String(); weekday == "Sunday" || weekday == "Saturday" {
		return true
	}
	return false
}

func IsLeapYear(input, format string) bool {
	date, dateFormat := getDateType(input, format)
	year, _, _ := returnYearMonthDay(date, dateFormat)
	return isLeapYear(year)
}

func isLeapYear(year int) bool {
	return ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0)
}

func GreaterThan(firstDate, secondDate Nakamura, format string) bool {
	date1, dateFormat1 := getDateType(firstDate.date, format)
	year1, month1, day1 := returnYearMonthDay(date1, dateFormat1)
	date2, dateFormat2 := getDateType(secondDate.date, format)
	year2, month2, day2 := returnYearMonthDay(date2, dateFormat2)
	return time.Date(year1, time.Month(month1), day1, 0, 0, 0, 0, time.UTC).After(time.Date(year2, time.Month(month2), day2, 0, 0, 0, 0, time.UTC))
}

func LessThan(firstDate, secondDate Nakamura, format string) bool {
	return !GreaterThan(firstDate, secondDate, format) && !Equal(firstDate, secondDate, format)
}

// Add adds value to the given unit component (Year, Month, or Day) of input.
// Add returns an error for an unrecognised Unit value.
func Add(input Nakamura, value int, unit Unit) (Nakamura, error) {
	date, dateFormat := getDateType(input.date, "YYYY-MM-DD")
	year, month, day := returnYearMonthDay(date, dateFormat)
	inputDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	switch unit {
	case Year:
		return Nakamura{strings.Split(inputDate.AddDate(value, 0, 0).String(), " ")[0], input.format}, nil
	case Month:
		return Nakamura{strings.Split(inputDate.AddDate(0, value, 0).String(), " ")[0], input.format}, nil
	case Day:
		return Nakamura{strings.Split(inputDate.AddDate(0, 0, value).String(), " ")[0], input.format}, nil
	}
	return Nakamura{}, fmt.Errorf("nakamura: unknown unit %d", unit)
}

func Equal(firstDate, secondDate Nakamura, format string) bool {
	date1, dateFormat1 := getDateType(firstDate.date, format)
	year1, month1, day1 := returnYearMonthDay(date1, dateFormat1)
	date2, dateFormat2 := getDateType(secondDate.date, format)
	year2, month2, day2 := returnYearMonthDay(date2, dateFormat2)
	return time.Date(year1, time.Month(month1), day1, 0, 0, 0, 0, time.UTC).Equal(time.Date(year2, time.Month(month2), day2, 0, 0, 0, 0, time.UTC))
}

func Weekday(input, format string) string {
	date, dateFormat := getDateType(input, format)
	year, month, day := returnYearMonthDay(date, dateFormat)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Weekday().String()
}

func monthName(input, format string) string {
	date, dateFormat := getDateType(input, format)
	_, month, _ := returnYearMonthDay(date, dateFormat)
	return getMonth(month).String()
}

// DaysInMonth returns the number of days in a month
// For example, for January 2017: daysInMonth(2017, 1) ==> 31
func DaysInMonth(input, format string) (int, error) {
	date, dateFormat := getDateType(input, format)
	year, month, _ := returnYearMonthDay(date, dateFormat)

	return getDaysInMonth(year, month)
}

func getDaysInMonth(year, month int) (int, error) {
	if month > 0 && month <= 12 {
		if month == 2 {
			if isLeapYear(year) {
				return 29, nil
			}
			return 28, nil
		}

		// Force-use a zero index to accommodate
		month = month - 1

		// Months 1 - 7 (Odd -> 31 | Even -> 30)
		// Months 8 - 12 (Odd -> 30 | Even -> 31)
		return 31 - (month % 7 % 2), nil
	}

	return 0, errors.New("Invalid month")
}
