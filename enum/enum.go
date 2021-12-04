// Package enum provides handy functional programming -style
// functions that operate on enumerable data types.
package enum

import (
	"constraints"
)

// All returns true if all f(item) calls return true.
func All[T any](items []T, f func(T) bool) bool {
	for _, v := range items {
		if !f(v) {
			return false
		}
	}
	return true
}

// AllWithIndex returns true if all f(index, item) calls return true.
func AllWithIndex[T any](items []T, f func(int, T) bool) bool {
	for i, v := range items {
		if !f(i, v) {
			return false
		}
	}
	return true
}

// Any returns true if any f(item) call returns true.
func Any[T any](items []T, f func(T) bool) bool {
	for _, v := range items {
		if f(v) {
			return true
		}
	}
	return false
}

// AnyWithIndex returns true if any f(index, item) call returns true.
func AnyWithIndex[T any](items []T, f func(int, T) bool) bool {
	for i, v := range items {
		if f(i, v) {
			return true
		}
	}
	return false
}

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

// Dedup returns a slice where all consecutive duplicated
// elements are collapsed to a single element.
func Dedup[T comparable](values []T) []T {
	result := []T{}
	for i, v := range values {
		if i == 0 || v != values[i-1] {
			result = append(result, v)
		}
	}
	return result
}

// Each processes each item with the provided function.
func Each[T any](items []T, f func(item T)) {
	for _, item := range items {
		f(item)
	}
}

// EachWithIndex iterates over a slice and calls the provided
// function with its index and value.
func EachWithIndex[T any](items []T, f func(i int, value T)) {
	for i, value := range items {
		f(i, value)
	}
}

// FilterFunc takes a value and returns true if the
// value is to be kept.
type FilterFunc[T any] func(T) bool

// Filter filters a slice of values and keeps
// those for which f(value) returns true.
func Filter[T any](values []T, f FilterFunc[T]) []T {
	var result []T
	for _, v := range values {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterFuncWithIndex takes an index and value and returns true if the
// value is to be kept.
type FilterFuncWithIndex[T any] func(int, T) bool

// FilterWithIndex filters a slice of values and keeps
// those for which f(index, value) returns true.
func FilterWithIndex[T any](values []T, f FilterFuncWithIndex[T]) []T {
	var result []T
	for i, v := range values {
		if f(i, v) {
			result = append(result, v)
		}
	}
	return result
}

// Find returns the first element for which f(value) returns true.
// If no element is found, defValue is returned.
func Find[T any](values []T, f FilterFunc[T], defValue T) T {
	for _, v := range values {
		if f(v) {
			return v
		}
	}
	return defValue
}

// FindWithIndex returns the first element for which f(index, value) returns true.
// If no element is found, defValue is returned.
func FindWithIndex[T any](values []T, f FilterFuncWithIndex[T], defValue T) T {
	for i, v := range values {
		if f(i, v) {
			return v
		}
	}
	return defValue
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

// Reduce reduces a slice using an accumulator.
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

// RunesWithIndex iterates over a string and calls the provided
// function with its index and rune. This is because I couldn't
// figure out how to make WithIndex work with a string and its runes.
func RunesWithIndex(s string, f func(i int, value rune)) {
	for i, v := range s {
		f(i, v)
	}
}
