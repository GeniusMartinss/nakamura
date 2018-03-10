package nakamura

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Months struct {
	January,
	February,
	March,
	April,
	May,
	June,
	July,
	August,
	September,
	October,
	November,
	December int
}

func IsDateValid(input, format string) bool {
	if len(input) == 0 {
		return false
	}
	date, dateFormat := getDateType(input, format)
	dateIntegerCount := 0

	if len(dateFormat) != 3 || len(date) != 3 {
		return false
	}

	//for each element in dateHyphen and dateSlash check if it's a valid integer
	for i := 0; i < len(date); i++ {
		if _, err := strconv.Atoi(date[i]); err == nil {
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
	validityCheck := Months{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	validityCount := Months{}
	if isLeap(year) {
		validityCount = Months{31, 29, 31, 30, 31, 30, 31, 30, 31, 31, 30, 31}
	} else {
		validityCount = Months{31, 28, 31, 30, 31, 30, 31, 30, 31, 31, 30, 31}

	}
	monthVal := reflect.ValueOf(validityCheck)
	daysVal := reflect.ValueOf(validityCount)

	//Get Month check if day passed in is valid in the given month
	for i := 0; i < monthVal.NumField(); i++ {
		if month == monthVal.Field(i).Interface().(int) {
			maxDaysInMonth := daysVal.Field(i).Interface().(int)
			if day <= maxDaysInMonth && day > 0 {
				return true
			} else {
				return false
			}
		}
	}

	return false
}

func Normalise(input, format string) string {
	date, dateFormat := getDateType(input, format)
	year, month, day := returnYearMonthDay(date, dateFormat)
	return strings.Split(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).String(), " ")[0]

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
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String() + "," + strconv.Itoa(day) + " " + getMonth(month) + " " + strconv.Itoa(year)
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

func getMonth(month int) string {
	validityCheck := Months{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	monthVal := reflect.ValueOf(validityCheck)
	for i := 0; i < monthVal.NumField(); i++ {
		if month == monthVal.Field(i).Interface().(int) {
			return monthVal.Type().Field(i).Name
		}
	}
	return ""
}

func Today() string {
	return strings.Split(time.Now().String(), " ")[0]
}

func IsWeekend(input, format string) bool {
	date, dateFormat := getDateType(input, format)
	year, month, day := returnYearMonthDay(date, dateFormat)

	if weekday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String(); weekday == "Sunday" || weekday == "Saturday" {
		return true
	}
	return false
}

func IsLeapYear(input, format string) bool {
	date, dateFormat := getDateType(input, format)
	year, _, _ := returnYearMonthDay(date, dateFormat)
	return ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0)
}
func isLeap(year int) bool {
	return ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0)
}
func GreaterThan(firstDate, secondDate Nakamura, format string) bool {
	date1, dateFormat1 := getDateType(firstDate.date, format)
	year1, month1, day1 := returnYearMonthDay(date1, dateFormat1)
	date2, dateFormat2 := getDateType(secondDate.date, format)
	year2, month2, day2 := returnYearMonthDay(date2, dateFormat2)
	return time.Date(year1, time.Month(month1), day1, 0, 0, 0, 0, time.Local).After(time.Date(year2, time.Month(month2), day2, 0, 0, 0, 0, time.Local))
}

func LessThan(firstDate, secondDate Nakamura, format string) bool {
	return !GreaterThan(firstDate, secondDate, format)
}

func Add(input Nakamura, value int, format string) Nakamura {
	date, dateFormat := getDateType(input.date, "YYYY-MM-DD")
	year, month, day := returnYearMonthDay(date, dateFormat)
	inputDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	switch format {
	case "YYYY":
		return Nakamura{strings.Split(inputDate.AddDate(value, 0, 0).String(), " ")[0], input.format}
	case "MM":
		return Nakamura{strings.Split(inputDate.AddDate(0, value, 0).String(), " ")[0], input.format}
	case "DD":
		return Nakamura{strings.Split(inputDate.AddDate(0, 0, value).String(), " ")[0], input.format}
	}
	return Nakamura{}
}

func Equal(firstDate, secondDate Nakamura, format string) bool {
	date1, dateFormat1 := getDateType(firstDate.date, format)
	year1, month1, day1 := returnYearMonthDay(date1, dateFormat1)
	date2, dateFormat2 := getDateType(secondDate.date, format)
	year2, month2, day2 := returnYearMonthDay(date2, dateFormat2)
	return time.Date(year1, time.Month(month1), day1, 0, 0, 0, 0, time.Local).Equal(time.Date(year2, time.Month(month2), day2, 0, 0, 0, 0, time.Local))
}

func Weekday(input, format string) string {
	date, dateFormat := getDateType(input, format)
	year, month, day := returnYearMonthDay(date, dateFormat)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String()
}

func Month(input, format string) string {
	date, dateFormat := getDateType(input, format)
	_, month, _ := returnYearMonthDay(date, dateFormat)
	return getMonth(month)
}
