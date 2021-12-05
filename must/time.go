package must

import (
	"log"
	"time"
)

// Parse parses a formatted string and returns the time value it represents.
// It dies if there is an error.
func Parse(layout, value string) time.Time {
	v, err := time.Parse(layout, value)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

// ParseInLocation is like Parse but differs in two important ways.
// First, in the absence of time zone information, Parse interprets
// a time as UTC; ParseInLocation interprets the time as in the given
// location. Second, when given a zone offset or abbreviation, Parse
// tries to match it against the Local location; ParseInLocation uses
// the given location.
func ParseInLocation(layout, value string, loc *time.Location) time.Time {
	v, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
