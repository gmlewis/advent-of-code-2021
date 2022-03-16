// Package enum provides handy functional programming -style
// functions that operate on enumerable data types.
package enum

import (
	"golang.org/x/exp/constraints"
)

// Equals returns a function that checks if a value
// is equal to a given value.
func Equals[T comparable](value T) FilterFunc[T] {
	return func(v T) bool { return v == value }
}

// Identity returns the value passed to it.
func Identity[T any](value T) T { return value }

// Length returns the length of the provided slice.
func Length[T any](items []T) int { return len(items) }

// First returns the first item of the provided slice or its zero value.
func First[T any](items []T) (ret T) {
	if len(items) == 0 {
		return ret
	}
	return items[0]
}

// Longer returns the longer slice.
// If len(a)==len(b), then a is preferred.
func Longer[T any](a, b []T) []T {
	if len(a) >= len(b) {
		return a
	}
	return b
}

// Shorter returns the shorter slice.
// If len(a)==len(b), then b is preferred.
func Shorter[T any](a, b []T) []T {
	if len(a) < len(b) {
		return a
	}
	return b
}

// Number has the "+" operator.
type Number interface {
	constraints.Integer | constraints.Float
}

// Average returns the average of a slice of numbers.
func Average[T Number](values []T) (ret T) {
	for _, v := range values {
		ret += v
	}
	return ret / T(len(values))
}

// Sum sums up a slice of numbers.
func Sum[T Number](values []T) (ret T) {
	for _, v := range values {
		ret += v
	}
	return ret
}

// Product multiples a slice of numbers together.
func Product[T Number](values []T) (ret T) {
	for i, v := range values {
		if i == 0 {
			ret = v
			continue
		}
		ret *= v
	}
	return ret
}
