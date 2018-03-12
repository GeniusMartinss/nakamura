# Nakamura
A lightweight golang date library for parsing, validating, manipulating, and formatting dates
supporting both slash style and hyphen style dates
## Install
```
go get -u github.com/geniusmartinss/nakamura
```
## Usage
```
import "github.com/geniusmartinss/nakamura"

nakamura.NewDate("2018-02-12", "YYYY-MM-DD") //{2018-02-12, "YYYY-MM-DD"}
nakamura.NewDate("", "YYYY-MM-DD") //Returns the date for the current day

testDate := nakamura.NewDate("2018-02-12", "YYYY-MM-DD")
testDate.IsDateVaid() //true

badDate := nakamura.NewDate("2018-03-32", "YYYY-MM-DD")
badDate.Normalise() //2018-04-01

testDate := nakamura.NewDate("2018-03-32", "YYYY-MM-DD")
testDate.Humanise() //Friday,9 November 2018

testDate := NewDate("2018/03/09", "YYYY/MM/DD")
testDate.IsWeekend() //false

testDate := NewDate("2016-11-09", "YYYY-MM-DD")
testDate.IsLeapYear() //true

firstDate := NewDate("2017-11-09", "YYYY-MM-DD")
secondDate := NewDate("2016-11-09","YYYY-MM-DD")
firstDate.Equal(secondDate) //false
firstDate.GreaterThan(firstDate) //true
firstDate.LessThan(secondDate) //false

testDate := NewDate("2012-11-09")
testDate.Add(3,"YYYY") //{"2015-11-09", "YYYY-MM-DD"}
testDate.Add(3,"DD") //{"2012-11-12", "YYYY-MM-DD"}
testDate.Subtract(3,"MM") //{"2012-08-09", "YYYY-MM-DD"}
```

## Contributing
Just Make a pull request :) Consider adding support for date

# Author
https://twitter.com/geniusmartins

## License
Released under the <a href ="http://www.opensource.org/licenses/MIT">MIT License</a>