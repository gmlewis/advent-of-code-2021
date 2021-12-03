package must

import (
	"log"
	"strconv"
)

// Atoi converts a string to an int.
func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

// ParseInt parses an integer with the given base and size.
// It dies if there is an error.
func ParseInt(s string, base, size int) int {
	v, err := strconv.ParseInt(s, base, size)
	if err != nil {
		log.Fatal(err)
	}
	return int(v)
}
