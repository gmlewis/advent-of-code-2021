// Package enum provides handy functional programming -style
// functions that operate on enumerable data types.
package enum

import (
	"constraints"
)

// Equals returns a function that checks if a value
// is equal to a given value.
func Equals[T comparable](value T) func(T) bool {
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

// Member checks if elem exists within values.
func Member[T comparable](values []T, elem T) bool {
	for _, v := range values {
		if elem == v {
			return true
		}
	}
	return false
}

// Reduce reduces a slice using an accumulator.
func Reduce[S any, T any](values []S, acc T, f func(S, T) T) T {
	for _, v := range values {
		acc = f(v, acc)
	}
	return acc
}

// Scan applies the given function to each element,
// emits the result and uses the same result as the accumulator for the
// next computation. It uses the given acc as the starting value.
func Scan[T any](items []T, acc T, f func(a, b T) T) (ret []T) {
	for _, v := range items {
		nv := f(acc, v)
		ret = append(ret, nv)
		acc = nv
	}

	return ret
}

// Uniq removes all duplicated elements.
func Uniq[T comparable](items []T) (ret []T) {
	seen := map[T]struct{}{}
	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		ret = append(ret, item)
		seen[item] = struct{}{}
	}

	return ret
}

// Zip zips corresponding elements from slice of slices.
//
// The zipping finishes as soon as any slice ends.
func Zip[T any](lists [][]T) (ret [][]T) {
	n := len(lists)
	var i int
outer:
	for {
		t := make([]T, 0, n)
		for _, list := range lists {
			if i >= len(list) {
				break outer
			}
			t = append(t, list[i])
		}
		ret = append(ret, t)
		i++
	}

	return ret
}

// Zip2 zips corresponding elements from different types into a slice of structs.
//
// The zipping finishes as soon as either slice ends.
func Zip2[S any, T any, KV any](sList []S, tList []T, f func(S, T) KV) (ret []KV) {
	for i := 0; i < len(sList) && i < len(tList); i++ {
		ret = append(ret, f(sList[i], tList[i]))
	}

	return ret
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
