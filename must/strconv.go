package must

import (
	"strconv"
)

// Atoi converts a string to an int.
func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		fatal(err)
	}
	return v
}

// ParseBool parses a boolean.
// It dies if there is an error.
func ParseBool(str string) bool {
	v, err := strconv.ParseBool(str)
	if err != nil {
		fatal(err)
	}
	return v
}

// ParseComplex parses a complex number with the given bitSize.
// It dies if there is an error.
func ParseComplex(s string, bitSize int) complex128 {
	v, err := strconv.ParseComplex(s, bitSize)
	if err != nil {
		fatal(err)
	}
	return v
}

// ParseFloat parses a complex number with the given bitSize.
// It dies if there is an error.
func ParseFloat(s string, bitSize int) float64 {
	v, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		fatal(err)
	}
	return v
}

// ParseInt parses an integer with the given base and bitSize.
// It dies if there is an error.
func ParseInt(s string, base, bitSize int) int {
	v, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		fatal(err)
	}
	return int(v)
}

// ParseUint parses an unsigned integer with the given base and bitSize.
// It dies if there is an error.
func ParseUint(s string, base int, bitSize int) uint64 {
	v, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		fatal(err)
	}
	return v
}
