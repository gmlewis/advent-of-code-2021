// Package strfn provides curried functional versions of the functions
// provided by the "strings" package so they can be more easily
// used within "enum" and "maps" functions.
package strfn

import "strings"

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
func Compare(b string) func(a string) int {
	return func(a string) int { return strings.Compare(a, b) }
}

// Contains reports whether substr is within s.
func Contains(substr string) func(s string) bool {
	return func(s string) bool { return strings.Contains(s, substr) }
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(chars string) func(s string) bool {
	return func(s string) bool { return strings.ContainsAny(s, chars) }
}
