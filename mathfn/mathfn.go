// Package mathfn provides curried functional versions of the functions
// provided by the "math" package so they can be more easily
// used within "enum" and "maps" functions.
package mathfn

import "constraints"

// Number is a number.
type Number interface {
	constraints.Integer | constraints.Unsigned | constraints.Float
}

// Abs returns the absolute value of the given number.
func Abs[T Number](v T) T {
	if v < T(0) {
		return -v
	}
	return v
}
