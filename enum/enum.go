// Package enum provides handy functional programming -style
// functions that operate on enumerable data types.
package enum

import (
	"constraints"
)

// ChunkEvery takes a slice of values and chunks them n-at-a-time
// with the given step size. It discards any left-over items.
func ChunkEvery[T any](values []T, n, step int) [][]T {
	result := [][]T{}
	for i := 0; i+n-1 < len(values); i += step {
		chunk := make([]T, 0, n)
		for j := 0; j < n; j++ {
			chunk = append(chunk, values[i+j])
		}
		result = append(result, chunk)
	}
	return result
}

// FilterFunc takes a value and returns true if the
// value is to be kept.
type FilterFunc[T any] func(T) bool

// Filter filters a slice of values and keeps
// those for which f returns true.
func Filter[T any](values []T, f FilterFunc[T]) []T {
	var result []T
	for _, v := range values {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map maps a slice of values from one type to another
// using the provided func f.
func Map[S any, T any](slice []T, f func(a T) S) []S {
	result := make([]S, 0, len(slice))
	for _, v := range slice {
		result = append(result, f(v))
	}
	return result
}

// Reduce reduces slices using an accumulator.
func Reduce[S any, T any](values []S, acc T, f func(S, T) T) T {
	for _, v := range values {
		acc = f(v, acc)
	}
	return acc
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
	constraints.Integer | constraints.Unsigned | constraints.Float | constraints.Complex
}

// Sum sums up a slice of numbers.
func Sum[T Number](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

// WithIndex iterates over a slice and calls the provided
// function with its index and value.
func WithIndex[T any](items []T, f func(i int, value T)) {
	for i, value := range items {
		f(i, value)
	}
}

// RunesWithIndex iterates over a string and calls the provided
// function with its index and rune. This is because I couldn't
// figure out how to make WithIndex work with a string and its runes.
func RunesWithIndex(s string, f func(i int, value rune)) {
	for i, v := range s {
		f(i, v)
	}
}
