package nakamura

import (
	"strings"
	"testing"
)

func TestNakamura_IsDateValid(t *testing.T) {
	cases := []struct {
		input Nakamura
		want  bool
	}{
		{Nakamura{"2011-11-09", "YYYY-MM-DD"}, true},
		{Nakamura{"2011/11/09", "YYYY/MM/DD"}, true},
		{Nakamura{"2011-11-09", "YYYY/MM/DD"}, false},
		{Nakamura{"11-2011-09", "YYYY-MM-DD"}, false},
	}
	for _, c := range cases {
		got := c.input.IsDateValid()
		if got != c.want {
			t.Errorf("IsDateValid(%q) == %t, want %t", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_IsLeapYear(t *testing.T) {
	cases := []struct {
		input Nakamura
		want  bool
	}{
		{Nakamura{"2016-11-09", "YYYY-MM-DD"}, true},
		{Nakamura{"2017/11/25", "YYYY/MM/DD"}, false},
	}
	for _, c := range cases {
		got := c.input.IsLeapYear()
		if got != c.want {
			t.Errorf("IsLeapYear(%q) == %t, want %t", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_IsWeekEnd(t *testing.T) {
	cases := []struct {
		input Nakamura
		want  bool
	}{
		{Nakamura{"2018/03/09", "YYYY/MM/DD"}, false},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, true},
	}
	for _, c := range cases {
		got := c.input.IsWeekEnd()
		if got != c.want {
			t.Errorf("IsWeekEnd(%q) == %t, want %t", c.input.date, got, c.want)
		}
	}
}

func TestGreaterThan(t *testing.T) {
	cases := []struct {
		firstDate, SecondDate Nakamura
		want                  bool
	}{
		{Nakamura{"2018-03-09", "YYYY-MM-DD"}, Nakamura{"2018-03-10", "YYYY-MM-DD"}, false},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, Nakamura{"2018-03-10", "YYYY-MM-DD"}, false},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, Nakamura{"2018-03-09", "YYYY-MM-DD"}, true},
	}
	for _, c := range cases {
		got := c.firstDate.GreaterThan(c.SecondDate)
		if got != c.want {
			t.Errorf("GreaterThan(%q) == %t, want %t", c.firstDate.date, got, c.want)
		}
	}
}

func TestNakamura_LessThan(t *testing.T) {
	cases := []struct {
		firstDate, SecondDate Nakamura
		want                  bool
	}{
		{Nakamura{"2018-03-09", "YYYY-MM-DD"}, Nakamura{"2018-03-10", "YYYY-MM-DD"}, true},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, Nakamura{"2018-03-09", "YYYY-MM-DD"}, false},
	}
	for _, c := range cases {
		got := c.firstDate.LessThan(c.SecondDate)
		if got != c.want {
			t.Errorf("LessThan(%q) == %t, want %t", c.firstDate.date, got, c.want)
		}
	}
}

func TestNakamura_Between(t *testing.T) {
	cases := []struct {
		firstDate, secondDate, thirdDate Nakamura
		want                             bool
	}{
		{Nakamura{"2018-03-09", "YYYY-MM-DD"}, Nakamura{"2018-03-08", "YYYY-MM-DD"}, Nakamura{"2018-03-10", "YYYY-MM-DD"}, true},
		{Nakamura{"2018-03-12", "YYYY-MM-DD"}, Nakamura{"2018-03-09", "YYYY-MM-DD"}, Nakamura{"2018-03-10", "YYYY-MM-DD"}, false},
		{Nakamura{"2018-03-09", "YYYY-MM-DD"}, Nakamura{"2018-03-09", "YYYY-MM-DD"}, Nakamura{"2018-03-10", "YYYY-MM-DD"}, true},
	}
	for _, c := range cases {
		got := c.firstDate.Between(c.secondDate, c.thirdDate)
		if got != c.want {
			t.Errorf("Between(%q) == %t, want %t", c.firstDate.date, got, c.want)
		}
	}
}

func TestEqual(t *testing.T) {
	cases := []struct {
		firstDate, SecondDate Nakamura
		want                  bool
	}{
		{Nakamura{"2018-03-09", "YYYY-MM-DD"}, Nakamura{"2018-03-09", "YYYY-MM-DD"}, true},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, Nakamura{"2018-03-09", "YYYY-MM-DD"}, false},
	}
	for _, c := range cases {
		got := c.firstDate.Equal(c.SecondDate)
		if got != c.want {
			t.Errorf("Equal(%q) == %t, want %t", c.firstDate.date, got, c.want)
		}
	}
}

func TestNakamura_Add(t *testing.T) {
	cases := []struct {
		input Nakamura
		value int
		unit  Unit
		want  Nakamura
	}{
		{Nakamura{"2018-03-09", "YYYY-MM-DD"}, 2, Month, Nakamura{"2018-05-09", "YYYY-MM-DD"}},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, 3, Day, Nakamura{"2018-03-13", "YYYY-MM-DD"}},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, 5, Year, Nakamura{"2023-03-10", "YYYY-MM-DD"}},
	}
	for _, c := range cases {
		got, err := c.input.Add(c.value, c.unit)
		if err != nil {
			t.Errorf("Add(%q) returned unexpected error: %v", c.input.date, err)
			continue
		}
		if got != c.want {
			t.Errorf("Add(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_Add_UnknownUnit(t *testing.T) {
	input := Nakamura{"2018-03-09", "YYYY-MM-DD"}
	_, err := input.Add(1, Unit(99))
	if err == nil {
		t.Error("Add with unknown unit: expected error, got nil")
	}
}

func TestNakamura_Subtract(t *testing.T) {
	cases := []struct {
		input Nakamura
		value int
		unit  Unit
		want  Nakamura
	}{
		{Nakamura{"2018-03-09", "YYYY-MM-DD"}, 2, Month, Nakamura{"2018-01-09", "YYYY-MM-DD"}},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, 3, Day, Nakamura{"2018-03-07", "YYYY-MM-DD"}},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, 5, Year, Nakamura{"2013-03-10", "YYYY-MM-DD"}},
	}
	for _, c := range cases {
		got, err := c.input.Subtract(c.value, c.unit)
		if err != nil {
			t.Errorf("Subtract(%q) returned unexpected error: %v", c.input.date, err)
			continue
		}
		if got != c.want {
			t.Errorf("Subtract(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_Subtract_UnknownUnit(t *testing.T) {
	input := Nakamura{"2018-03-09", "YYYY-MM-DD"}
	_, err := input.Subtract(1, Unit(99))
	if err == nil {
		t.Error("Subtract with unknown unit: expected error, got nil")
	}
}

func TestNakamura_Weekday(t *testing.T) {
	cases := []struct {
		input Nakamura
		want  string
	}{
		{Nakamura{"2018-03-05", "YYYY-MM-DD"}, "Monday"},
		{Nakamura{"2018-03-06", "YYYY-MM-DD"}, "Tuesday"},
		{Nakamura{"2018-03-07", "YYYY-MM-DD"}, "Wednesday"},
		{Nakamura{"2018-03-08", "YYYY-MM-DD"}, "Thursday"},
		{Nakamura{"2018-03-09", "YYYY-MM-DD"}, "Friday"},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, "Saturday"},
		{Nakamura{"2018-03-11", "YYYY-MM-DD"}, "Sunday"},
	}
	for _, c := range cases {
		got := c.input.Weekday()
		if got != c.want {
			t.Errorf("Weekday(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_Month(t *testing.T) {
	cases := []struct {
		input Nakamura
		want  string
	}{
		{Nakamura{"2018-01-05", "YYYY-MM-DD"}, "January"},
		{Nakamura{"2018-02-06", "YYYY-MM-DD"}, "February"},
		{Nakamura{"2018-03-07", "YYYY-MM-DD"}, "March"},
		{Nakamura{"2018-04-08", "YYYY-MM-DD"}, "April"},
		{Nakamura{"2018-05-09", "YYYY-MM-DD"}, "May"},
		{Nakamura{"2018-06-10", "YYYY-MM-DD"}, "June"},
		{Nakamura{"2018-07-11", "YYYY-MM-DD"}, "July"},
		{Nakamura{"2018-08-11", "YYYY-MM-DD"}, "August"},
		{Nakamura{"2018-09-11", "YYYY-MM-DD"}, "September"},
		{Nakamura{"2018-10-11", "YYYY-MM-DD"}, "October"},
		{Nakamura{"2018-11-11", "YYYY-MM-DD"}, "November"},
		{Nakamura{"2018-12-11", "YYYY-MM-DD"}, "December"},
	}
	for _, c := range cases {
		got := c.input.Month()
		if got != c.want {
			t.Errorf("Month(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_Normalise(t *testing.T) {
	cases := []struct {
		input Nakamura
		want  Nakamura
	}{
		{Nakamura{"2018-13-09", "YYYY-MM-DD"}, Nakamura{"2019-01-09", "YYYY-MM-DD"}},
		{Nakamura{"2018-03-32", "YYYY-MM-DD"}, Nakamura{"2018-04-01", "YYYY-MM-DD"}},
	}
	for _, c := range cases {
		got := c.input.Normalise()
		if got != c.want {
			t.Errorf("Normalise(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestHumanise(t *testing.T) {
	cases := []struct {
		input Nakamura
		want  string
	}{
		{Nakamura{"2018-11-09", "YYYY-MM-DD"}, "Friday,9 November 2018"},
		{Nakamura{"2018-12-10", "YYYY-MM-DD"}, "Monday,10 December 2018"},
	}
	for _, c := range cases {
		got := c.input.Humanise()
		if got != c.want {
			t.Errorf("Humanise(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_IsFuture(t *testing.T) {
	cases := []struct {
		input Nakamura
		want  bool
	}{
		{Nakamura{"1998-11-09", "YYYY-MM-DD"}, false},
		//Jesus woulda been back by then :)
		{Nakamura{"9000-12-10", "YYYY-MM-DD"}, true},
	}
	for _, c := range cases {
		got := c.input.IsFuture()
		if got != c.want {
			t.Errorf("IsFuture(%q) == %t, want %t", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_IsPast(t *testing.T) {
	cases := []struct {
		input Nakamura
		want  bool
	}{
		{Nakamura{"1998-11-09", "YYYY-MM-DD"}, true},
		//Jesus woulda been back by then :)
		{Nakamura{"9000-12-10", "YYYY-MM-DD"}, false},
	}
	for _, c := range cases {
		got := c.input.IsPast()
		if got != c.want {
			t.Errorf("IsFuture(%q) == %t, want %t", c.input.date, got, c.want)
		}
	}
}

func TestDaysInMonth(t *testing.T) {
	cases := []struct {
		input Nakamura
		want  int
	}{
		{Nakamura{"2018-01-10", "YYYY-MM-DD"}, 31},
		{Nakamura{"2018-02-10", "YYYY-MM-DD"}, 28},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, 31},
		{Nakamura{"2018-04-10", "YYYY-MM-DD"}, 30},
		{Nakamura{"2018-05-10", "YYYY-MM-DD"}, 31},
		{Nakamura{"2018-06-10", "YYYY-MM-DD"}, 30},
		{Nakamura{"2018-07-10", "YYYY-MM-DD"}, 31},
		{Nakamura{"2018-08-10", "YYYY-MM-DD"}, 31},
		{Nakamura{"2018-09-10", "YYYY-MM-DD"}, 30},
		{Nakamura{"2018-10-10", "YYYY-MM-DD"}, 31},
		{Nakamura{"2018-11-10", "YYYY-MM-DD"}, 30},
		{Nakamura{"2018-12-10", "YYYY-MM-DD"}, 31},
	}

	errorCases := []struct {
		input Nakamura
	}{
		{Nakamura{"2018-13-10", "YYYY-MM-DD"}},
	}

	for _, c := range cases {
		got, _ := c.input.MonthDays()
		if got != c.want {
			t.Errorf("MonthDays(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}

	for _, c := range errorCases {
		_, got := c.input.MonthDays()
		_, ok := got.(error)
		if !ok {
			t.Errorf("MonthDays(%q) == %q, want instance of error", c.input.date, got)
		}
	}

}

func TestMax(t *testing.T) {
	cases := []struct {
		args []Nakamura
		want Nakamura
	}{
		{
			[]Nakamura{
				{"2018-03-09", "YYYY-MM-DD"},
				{"2018-03-10", "YYYY-MM-DD"},
				{"2018-03-11", "YYYY-MM-DD"},
			},
			Nakamura{"2018-03-11", "YYYY-MM-DD"},
		},
	}
	for _, c := range cases {
		got := Max((c.args)...)
		if got != c.want {
			t.Errorf("Max(%v) == %v, want %v", c.args, got, c.want)
		}
	}
}

func TestMin(t *testing.T) {
	cases := []struct {
		args []Nakamura
		want Nakamura
	}{
		{
			[]Nakamura{
				{"2018-03-09", "YYYY-MM-DD"},
				{"2018-03-10", "YYYY-MM-DD"},
				{"2018-03-11", "YYYY-MM-DD"},
			},
			Nakamura{"2018-03-09", "YYYY-MM-DD"},
		},
	}
	for _, c := range cases {
		got := Min((c.args)...)
		if got != c.want {
			t.Errorf("Min(%v) == %v, want %v", c.args, got, c.want)
		}
	}
}

// TestNewDate exercises all three behavioural rules of NewDate:
//  1. empty / whitespace-only input → returns today's date with nil error
//  2. datetime string (contains space or 'T') → returns Nakamura{} and a descriptive error
//  3. pure date string → returns Nakamura{date, format} with nil error
func TestNewDate(t *testing.T) {
	// ── rule 1: empty / whitespace input → today, nil error ──────────────────
	t.Run("empty string returns today with nil error", func(t *testing.T) {
		got, err := NewDate("", "YYYY-MM-DD")
		if err != nil {
			t.Fatalf("NewDate(\"\", ...) unexpected error: %v", err)
		}
		today := Today()
		if got.date != today {
			t.Errorf("NewDate(\"\", ...) date = %q, want today %q", got.date, today)
		}
		if got.format != "YYYY-MM-DD" {
			t.Errorf("NewDate(\"\", ...) format = %q, want %q", got.format, "YYYY-MM-DD")
		}
	})

	t.Run("whitespace-only string returns today with nil error", func(t *testing.T) {
		got, err := NewDate("   ", "YYYY-MM-DD")
		if err != nil {
			t.Fatalf("NewDate(\"   \", ...) unexpected error: %v", err)
		}
		today := Today()
		if got.date != today {
			t.Errorf("NewDate(\"   \", ...) date = %q, want today %q", got.date, today)
		}
	})

	// ── rule 2: datetime strings are rejected ─────────────────────────────────
	datetimeCases := []struct {
		name  string
		input string
	}{
		{"datetime with space separator", "2024-01-15 10:30:00"},
		{"ISO 8601 datetime with T separator", "2024-01-15T10:30:00"},
		{"date with time zone offset via T", "2024-01-15T00:00:00Z"},
	}
	for _, tc := range datetimeCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewDate(tc.input, "YYYY-MM-DD")
			if err == nil {
				t.Errorf("NewDate(%q) expected an error, got nil (result: %+v)", tc.input, got)
			}
			if (got != Nakamura{}) {
				t.Errorf("NewDate(%q) expected empty Nakamura{}, got %+v", tc.input, got)
			}
			if !strings.Contains(err.Error(), tc.input) {
				t.Errorf("NewDate(%q) error message %q should contain the offending input", tc.input, err.Error())
			}
		})
	}

	// ── rule 3: pure date string → Nakamura{date, format} with nil error ──────
	pureDateCases := []struct {
		date   string
		format string
	}{
		{"2024-01-15", "YYYY-MM-DD"},
		{"2024/06/30", "YYYY/MM/DD"},
		{"1999-12-31", "YYYY-MM-DD"},
	}
	for _, tc := range pureDateCases {
		tc := tc
		t.Run("pure date "+tc.date, func(t *testing.T) {
			got, err := NewDate(tc.date, tc.format)
			if err != nil {
				t.Fatalf("NewDate(%q, %q) unexpected error: %v", tc.date, tc.format, err)
			}
			if got.date != tc.date {
				t.Errorf("NewDate(%q, %q) .date = %q, want %q", tc.date, tc.format, got.date, tc.date)
			}
			if got.format != tc.format {
				t.Errorf("NewDate(%q, %q) .format = %q, want %q", tc.date, tc.format, got.format, tc.format)
			}
		})
	}
}
