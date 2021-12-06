package must

import (
	"time"
)

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of decimal numbers,
// each with optional fraction and a unit suffix, such as "300ms",
// "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func ParseDuration(s string) time.Duration {
	v, err := time.ParseDuration(s)
	if err != nil {
		fatal(err)
	}
	return v
}

// ParseTime parses a formatted string and returns the time value it represents.
// It dies if there is an error.
func ParseTime(layout, value string) time.Time {
	v, err := time.Parse(layout, value)
	if err != nil {
		fatal(err)
	}
	return v
}

// ParseTimeInLocation is like Parse but differs in two important ways.
// First, in the absence of time zone information, Parse interprets
// a time as UTC; ParseInLocation interprets the time as in the given
// location. Second, when given a zone offset or abbreviation, Parse
// tries to match it against the Local location; ParseInLocation uses
// the given location.
func ParseTimeInLocation(layout, value string, loc *time.Location) time.Time {
	v, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		fatal(err)
	}
	return v
}
