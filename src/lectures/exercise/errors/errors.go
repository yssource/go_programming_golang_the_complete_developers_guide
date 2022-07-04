//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour, Minute, Second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%s, %s", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	vals := strings.Split(input, ":")

	if len(vals) != 3 {
		return Time{}, &TimeParseError{
			msg:   "invalid number of time component",
			input: input,
		}
	}

	hour, err := strconv.Atoi(vals[0])
	if err != nil {
		return Time{}, &TimeParseError{
			msg:   fmt.Sprintf("error parsing hour: %v", err),
			input: input,
		}
	}
	if hour < 0 || hour > 23 {
		return Time{}, &TimeParseError{
			msg:   fmt.Sprintf("hour out of range: should be in [0, 24) range"),
			input: fmt.Sprintf("%d", hour),
		}
	}

	minute, err := strconv.Atoi(vals[1])
	if err != nil {
		return Time{}, &TimeParseError{
			msg:   fmt.Sprintf("error parsing minute: %v", err),
			input: input,
		}
	}
	if minute < 0 || minute > 59 {
		return Time{}, &TimeParseError{
			msg:   fmt.Sprintf("hour out of range: should be in [0, 60) range"),
			input: fmt.Sprintf("%d", minute),
		}
	}

	second, err := strconv.Atoi(vals[2])
	if err != nil {
		return Time{}, &TimeParseError{
			msg:   fmt.Sprintf("error parsing second: %v", err),
			input: input,
		}
	}
	if second < 0 || second > 59 {
		return Time{}, &TimeParseError{
			msg:   fmt.Sprintf("second out of range: should be in [0, 60) range"),
			input: fmt.Sprintf("%d", second),
		}
	}

	return Time{Hour: hour, Minute: minute, Second: second}, nil
}
