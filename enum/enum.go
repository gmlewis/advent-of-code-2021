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

// Count returns the count of items in the slice for which `f(item)` returns true.
func Count[T any](items []T, f func(T) bool) int {
	var result int
	for _, v := range items {
		if f(v) {
			result++
		}
	}
	return result
}

// CountWithIndex returns the count of items in the slice for which
// `f(index, item)` returns true.
func CountWithIndex[T any](items []T, f func(int, T) bool) int {
	var result int
	for i, v := range items {
		if f(i, v) {
			result++
		}
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

// FilterMap filters a slice of mapped values and keeps
// those for which f(value) returns true.
func FilterMap[S any, T any](values []S, m func(S) T, f FilterFunc[T]) []S {
	var result []S
	for _, v := range values {
		if f(m(v)) {
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

// FlatMap maps the given f(value) and flattens the result.
//
// For example:
//   FlatMap([]int{1,2,3}, func (v int) []string {
//     s := fmt.Sprintf("%v", v)
//     return []string{s,s}
//   })
// returns:
//   []string{"1","1","2","2","3","3"}
func FlatMap[S any, T any](values []S, f func(S) []T) []T {
	result := []T{}
	for _, v := range values {
		result = append(result, f(v)...)
	}
	return result
}

// FlatMapWithIndex maps the given f(index, value) and flattens the result.
func FlatMapWithIndex[S any, T any](values []S, f func(int, S) []T) []T {
	result := []T{}
	for i, v := range values {
		result = append(result, f(i, v)...)
	}
	return result
}

// Frequencies returns a map with keys as unique elements of the
// provided items and the values as the count of every item.
func Frequencies[T comparable](items []T) map[T]int {
	result := map[T]int{}
	for _, item := range items {
		result[item]++
	}
	return result
}

// FrequenciesBy returns a map with keys as unique elements of
// keyFunc(item) and the values as the count of every item.
func FrequenciesBy[S any, T comparable](items []S, keyFunc func(S) T) map[T]int {
	result := map[T]int{}
	for _, item := range items {
		result[keyFunc(item)]++
	}
	return result
}

// GroupBy splits the items into groups based on keyFunc and valueFunc.
//
// For example:
//    GroupBy([]string{"ant", "buffalo", "cat", "dingo"}, StrLength, Identity[string])
// returns:
//    {3: {"ant", "cat"}, 5: {"dingo"}, 7: {"buffalo"}}
// and
//    GroupBy([]string{ant buffalo cat dingo}, StrLength, StrFirst)
// returns:
//    {3: {"a", "c"}, 5: {"d"}, 7: {"b"}}
func GroupBy[K comparable, V any, T any](items []T, keyFunc func(T) K, valueFunc func(T) V) map[K][]V {
	result := map[K][]V{}
	for _, item := range items {
		k := keyFunc(item)
		result[k] = append(result[k], valueFunc(item))
	}
	return result
}

// Map maps a slice of values from one type to another
// using the provided func f.
func Map[S any, T any](values []T, f func(a T) S) []S {
	result := make([]S, 0, len(values))
	for _, v := range values {
		result = append(result, f(v))
	}
	return result
}

// Max returns the maximal element in the slice
// (or the zero value for an empty slice).
func Max[T constraints.Ordered](values []T) (ret T) {
	for i, v := range values {
		if i == 0 || v > ret {
			ret = v
		}
	}
	return ret
}

// Min returns the minimal element in the slice
// (or the zero value for an empty slice).
func Min[T constraints.Ordered](values []T) (ret T) {
	for i, v := range values {
		if i == 0 || v < ret {
			ret = v
		}
	}
	return ret
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
