package nakamura

import "testing"

func TestNakamura_IsDateValid(t *testing.T) {
	cases := [] struct{
		input Nakamura
		want bool
	}{
		{Nakamura{"2011-11-09","YYYY-MM-DD"}, true},
		{Nakamura{"2011/11/09","YYYY/MM/DD"}, true},
		{Nakamura{"2011-11-09","YYYY/MM/DD"}, false},
		{Nakamura{"11-2011-09","YYYY-MM-DD"}, false},
	}
	for _,c := range cases{
		got := c.input.IsDateValid()
		if got != c.want{
			t.Errorf("IsDateValid(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_IsLeapYear(t *testing.T) {
	cases := [] struct{
		input Nakamura
		want bool
	}{
		{Nakamura{"2016-11-09","YYYY-MM-DD"}, true},
		{Nakamura{"2017/11/25","YYYY/MM/DD"}, false},
	}
	for _,c := range cases{
		got := c.input.IsLeapYear()
		if got != c.want{
			t.Errorf("IsLeapYear(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_IsWeekEnd(t *testing.T) {
	cases := [] struct{
		input Nakamura
		want bool
	}{
		{Nakamura{"2018/03/09","YYYY/MM/DD"}, false},
		{Nakamura{"2018-03-10","YYYY-MM-DD"}, true},
	}
	for _,c := range cases{
		got := c.input.IsWeekEnd()
		if got != c.want{
			t.Errorf("IsWeekEnd(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestGreaterThan(t *testing.T) {
	cases := [] struct{
		firstDate, SecondDate Nakamura
		want bool
	}{
		{Nakamura{"2018-03-09", "YYYY-MM-DD"},Nakamura{"2018-03-10", "YYYY-MM-DD"}, false},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"},Nakamura{"2018-03-09", "YYYY-MM-DD"}, true},
	}
	for _,c := range cases{
		got := c.firstDate.GreaterThan(c.SecondDate)
		if got != c.want{
			t.Errorf("GreaterThan(%q) == %q, want %q", c.firstDate.date, got, c.want)
		}
	}
}

func TestNakamura_LessThan(t *testing.T) {
	cases := [] struct{
		firstDate, SecondDate Nakamura
		want bool
	}{
		{Nakamura{"2018-03-09", "YYYY-MM-DD"},Nakamura{"2018-03-10", "YYYY-MM-DD"},  true},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"},Nakamura{"2018-03-09", "YYYY-MM-DD"},  false},
	}
	for _,c := range cases{
		got := c.firstDate.LessThan(c.SecondDate)
		if got != c.want{
			t.Errorf("LessThan(%q) == %q, want %q", c.firstDate.date, got, c.want)
		}
	}
}

func TestEqual(t *testing.T) {
	cases := [] struct{
		firstDate, SecondDate Nakamura
		want bool
	}{
		{Nakamura{"2018-03-09", "YYYY-MM-DD"},Nakamura{"2018-03-09", "YYYY-MM-DD"},  true},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"},Nakamura{"2018-03-09", "YYYY-MM-DD"}, false},
	}
	for _,c := range cases{
		got := c.firstDate.Equal(c.SecondDate)
		if got != c.want{
			t.Errorf("Equal(%q) == %q, want %q", c.firstDate.date, got, c.want)
		}
	}
}

func TestNakamura_Add(t *testing.T) {
	cases := [] struct{
		input Nakamura
		value int
		format string
		want Nakamura
	}{
		{Nakamura{"2018-03-09", "YYYY-MM-DD"},2, "MM", Nakamura{"2018-05-09", "YYYY-MM-DD"}},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"},3, "DD", Nakamura{"2018-03-13", "YYYY-MM-DD"}},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"},5, "YYYY", Nakamura{"2023-03-10", "YYYY-MM-DD"}},
	}
	for _,c := range cases{
		got := c.input.Add(c.value, c.format)
		if got != c.want{
			t.Errorf("Add(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_Subtract(t *testing.T) {
	cases := [] struct{
		input Nakamura
		value int
		format string
		want Nakamura
	}{
		{Nakamura{"2018-03-09","YYYY-MM-DD"},2, "MM", Nakamura{"2018-01-09","YYYY-MM-DD"}},
		{Nakamura{"2018-03-10","YYYY-MM-DD"},3, "DD", Nakamura{"2018-03-07","YYYY-MM-DD"}},
		{Nakamura{"2018-03-10","YYYY-MM-DD"},5, "YYYY", Nakamura{"2013-03-10","YYYY-MM-DD"}},
	}
	for _,c := range cases{
		got := c.input.Subtract(c.value, c.format)
		if got != c.want{
			t.Errorf("Subtract(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_Weekday(t *testing.T) {
	cases := [] struct{
		input Nakamura
		want string
	}{
		{Nakamura{"2018-03-05", "YYYY-MM-DD"}, "Monday"},
		{Nakamura{"2018-03-06", "YYYY-MM-DD"},  "Tuesday"},
		{Nakamura{"2018-03-07", "YYYY-MM-DD"},  "Wednesday"},
		{Nakamura{"2018-03-08", "YYYY-MM-DD"},  "Thursday"},
		{Nakamura{"2018-03-09", "YYYY-MM-DD"},  "Friday"},
		{Nakamura{"2018-03-10", "YYYY-MM-DD"}, "Saturday"},
		{Nakamura{"2018-03-11", "YYYY-MM-DD"}, "Sunday"},
	}
	for _,c := range cases{
		got := c.input.Weekday()
		if got != c.want{
			t.Errorf("Weekday(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_Month(t *testing.T) {
	cases := [] struct{
		input Nakamura
		want string
	}{
		{Nakamura{"2018-01-05", "YYYY-MM-DD"},  "January"},
		{Nakamura{"2018-02-06", "YYYY-MM-DD"},  "February"},
		{Nakamura{"2018-03-07", "YYYY-MM-DD"}, "March"},
		{Nakamura{"2018-04-08", "YYYY-MM-DD"},  "April"},
		{Nakamura{"2018-05-09", "YYYY-MM-DD"}, "May"},
		{Nakamura{"2018-06-10", "YYYY-MM-DD"}, "June"},
		{Nakamura{"2018-07-11", "YYYY-MM-DD"}, "July"},
		{Nakamura{"2018-08-11", "YYYY-MM-DD"}, "August"},
		{Nakamura{"2018-09-11", "YYYY-MM-DD"}, "September"},
		{Nakamura{"2018-10-11", "YYYY-MM-DD"}, "October"},
		{Nakamura{"2018-11-11", "YYYY-MM-DD"}, "November"},
		{Nakamura{"2018-12-11", "YYYY-MM-DD"}, "December"},

	}
	for _,c := range cases{
		got := c.input.Month()
		if got != c.want{
			t.Errorf("Month(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestNakamura_Normalise(t *testing.T) {
	cases := [] struct{
		input Nakamura
		want Nakamura
	}{
		{Nakamura{"2018-13-09","YYYY-MM-DD"}, Nakamura{"2019-01-09","YYYY-MM-DD"}},
		{Nakamura{"2018-03-32","YYYY-MM-DD"}, Nakamura{"2018-04-01","YYYY-MM-DD"}},
	}
	for _,c := range cases{
		got := c.input.Normalise()
		if got != c.want{
			t.Errorf("Normalise(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}

func TestHumanise(t *testing.T) {
	cases := [] struct{
		input Nakamura
		want string
	}{
		{Nakamura{"2018-11-09","YYYY-MM-DD"}, "Friday,9 November 2018"},
		{Nakamura{"2018-12-10","YYYY-MM-DD"},"Monday,10 December 2018"},
	}
	for _,c := range cases{
		got := c.input.Humanise()
		if got != c.want{
			t.Errorf("Humanise(%q) == %q, want %q", c.input.date, got, c.want)
		}
	}
}