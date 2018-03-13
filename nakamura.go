package nakamura

import "strings"

type Nakamura struct {
	date, format string
}

func NewDate(date, format string) Nakamura {

	if len(strings.TrimSpace(date)) == 0 {
		return Nakamura{Today(), "YYYY-MM-DD"}
	}
	return Nakamura{strings.Split(date, " ")[0], format}
}

func (date Nakamura) IsDateValid() bool {
	return IsDateValid(date.date, date.format)
}

func (date Nakamura) Normalise() Nakamura {
	return Nakamura{Normalise(date.date, date.format), date.format}
}

func (date Nakamura) Humanise() string {
	return Humanise(date.date, date.format)
}

func (date Nakamura) IsWeekEnd() bool {
	return IsWeekend(date.date, date.format)
}

func (date Nakamura) IsLeapYear() bool {
	return IsLeapYear(date.date, date.format)
}

func (firstDate Nakamura) GreaterThan(secondDate Nakamura) bool {
	return GreaterThan(firstDate, secondDate, secondDate.format)
}

func (firstDate Nakamura) LessThan(secondDate Nakamura) bool {
	return LessThan(firstDate, secondDate, secondDate.format)
}

func (date Nakamura) Add(value int, format string) Nakamura {
	return Add(date, value, format)
}

func (date Nakamura) Subtract(value int, format string) Nakamura {
	return Add(date, -value, format)
}

func (firstDate Nakamura) Equal(secondDate Nakamura) bool {
	return Equal(firstDate, secondDate, secondDate.format)
}

func (date Nakamura) Weekday() string {
	return Weekday(date.date, date.format)
}

func (date Nakamura) Month() string {
	return Month(date.date, date.format)
}

func (date Nakamura) MonthDays() (int, error) {
	return DaysInMonth(date.date, date.format)
}

func Max(dates ...Nakamura) Nakamura {
	return GetMax(dates...)
}

func Min(dates ...Nakamura) Nakamura {
	return GetMin(dates...)
}
