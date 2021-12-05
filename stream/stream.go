// Package stream provides handy functional programming -style
// functions that operate on enumerable data types using channels.
package stream

import (
	"constraints"
	"log"
)

// First returns the first item of the provided chan or its zero value.
func First[T any](ch <-chan T) (ret T) {
	return <-ch
}

// Take returns the first n items from the provided chan.
func Take[T any](ch <-chan T, n int) []T {
	ret := make([]T, 0, n)
	for i := 0; i < n; i++ {
		ret = append(ret, <-ch)
	}
	return ret
}

// ToSlice converts the channel values to a slice.
func ToSlice[T any](ch <-chan T) []T {
	var ret []T
	for v := range ch {
		ret = append(ret, v)
	}
	return ret
}

// All returns true if all f(item) calls return true.
func All[T any](ch <-chan T, f func(T) bool) bool {
	for v := range ch {
		if !f(v) {
			return false
		}
	}
	return true
}

// AllWithIndex returns true if all f(index, item) calls return true.
func AllWithIndex[T any](ch <-chan T, f func(int, T) bool) bool {
	var i int
	for v := range ch {
		if !f(i, v) {
			return false
		}
		i++
	}
	return true
}

// Any returns true if any f(item) call returns true.
func Any[T any](ch <-chan T, f func(T) bool) bool {
	for v := range ch {
		if f(v) {
			return true
		}
	}
	return false
}

// AnyWithIndex returns true if any f(index, item) call returns true.
func AnyWithIndex[T any](ch <-chan T, f func(int, T) bool) bool {
	var i int
	for v := range ch {
		if f(i, v) {
			return true
		}
		i++
	}
	return false
}

// Count returns the count of items in the channel for which `f(item)` returns true.
func Count[T any](ch <-chan T, f func(T) bool) int {
	var result int
	for v := range ch {
		if f(v) {
			result++
		}
	}
	return result
}

// ChunkEvery takes a channel of values and chunks them n-at-a-time
// with the given step size. It discards any left-over items.
func ChunkEvery[T any](ch <-chan T, n, step int) <-chan []T {
	if step > n {
		log.Fatalf("step(%v) must be <= n(%v)", step, n)
	}

	outCh := make(chan []T, defaultBufSize)

	go func() {
		chunk := make([]T, 0, n)
		for v := range ch {
			chunk = append(chunk, v)
			if len(chunk) == n {
				outCh <- chunk
				chunk = chunk[step:]
			}
		}
		close(outCh)
	}()

	return outCh
}

// CountWithIndex returns the count of items in the channel for which
// `f(index, item)` returns true.
func CountWithIndex[T any](ch <-chan T, f func(int, T) bool) int {
	var i, result int
	for v := range ch {
		if f(i, v) {
			result++
		}
		i++
	}
	return result
}

// Dedup returns a channel where all consecutive duplicated
// elements are collapsed to a single element.
func Dedup[T comparable](ch <-chan T) []T {
	result := []T{}
	for v := range ch {
		if len(result) == 0 || v != result[len(result)-1] {
			result = append(result, v)
		}
	}
	return result
}

// Each processes each item with the provided function.
func Each[T any](ch <-chan T, f func(item T)) {
	for item := range ch {
		f(item)
	}
}

// EachWithIndex iterates over a channel and calls the provided
// function with its index and value.
func EachWithIndex[T any](ch <-chan T, f func(i int, value T)) {
	var i int
	for value := range ch {
		f(i, value)
		i++
	}
}

// FilterFunc takes a value and returns true if the
// value is to be kept.
type FilterFunc[T any] func(T) bool

// Filter filters a channel of values and keeps
// those for which f(value) returns true.
func Filter[T any](ch <-chan T, f FilterFunc[T]) []T {
	var result []T
	for v := range ch {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterMap filters a channel of mapped values and keeps
// those for which f(value) returns true.
func FilterMap[S any, T any](ch <-chan S, m func(S) T, f FilterFunc[T]) []S {
	var result []S
	for v := range ch {
		if f(m(v)) {
			result = append(result, v)
		}
	}
	return result
}

// FilterFuncWithIndex takes an index and value and returns true if the
// value is to be kept.
type FilterFuncWithIndex[T any] func(int, T) bool

// FilterWithIndex filters a channel of values and keeps
// those for which f(index, value) returns true.
func FilterWithIndex[T any](ch <-chan T, f FilterFuncWithIndex[T]) []T {
	var result []T
	var i int
	for v := range ch {
		if f(i, v) {
			result = append(result, v)
		}
		i++
	}
	return result
}

// Find returns the first element for which f(value) returns true.
// If no element is found, defValue is returned.
func Find[T any](ch <-chan T, f FilterFunc[T], defValue T) T {
	for v := range ch {
		if f(v) {
			return v
		}
	}
	return defValue
}

// FindWithIndex returns the first element for which f(index, value) returns true.
// If no element is found, defValue is returned.
func FindWithIndex[T any](ch <-chan T, f FilterFuncWithIndex[T], defValue T) T {
	var i int
	for v := range ch {
		if f(i, v) {
			return v
		}
		i++
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
func FlatMap[S any, T any](ch <-chan S, f func(S) []T) []T {
	result := []T{}
	for v := range ch {
		result = append(result, f(v)...)
	}
	return result
}

// FlatMapWithIndex maps the given f(index, value) and flattens the result.
func FlatMapWithIndex[S any, T any](ch <-chan S, f func(int, S) []T) []T {
	result := []T{}
	var i int
	for v := range ch {
		result = append(result, f(i, v)...)
		i++
	}
	return result
}

// Frequencies returns a map with keys as unique elements of the
// provided items and the values as the count of every item.
func Frequencies[T comparable](ch <-chan T) map[T]int {
	result := map[T]int{}
	for item := range ch {
		result[item]++
	}
	return result
}

// FrequenciesBy returns a map with keys as unique elements of
// keyFunc(item) and the values as the count of every item.
func FrequenciesBy[S any, T comparable](ch <-chan S, keyFunc func(S) T) map[T]int {
	result := map[T]int{}
	for item := range ch {
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
func GroupBy[K comparable, V any, T any](ch <-chan T, keyFunc func(T) K, valueFunc func(T) V) map[K][]V {
	result := map[K][]V{}
	for item := range ch {
		k := keyFunc(item)
		result[k] = append(result[k], valueFunc(item))
	}
	return result
}

// Map maps a channel of values from one type to another
// using the provided func f.
func Map[S any, T any](ch <-chan T, f func(a T) S) []S {
	ret := []S{}
	for v := range ch {
		ret = append(ret, f(v))
	}
	return ret
}

// Max returns the maximal element in the channel
// (or the zero value for an empty channel).
func Max[T constraints.Ordered](ch <-chan T) (ret T) {
	var i int
	for v := range ch {
		if i == 0 || v > ret {
			ret = v
		}
		i++
	}
	return ret
}

// Min returns the minimal element in the channel
// (or the zero value for an empty channel).
func Min[T constraints.Ordered](ch <-chan T) (ret T) {
	var i int
	for v := range ch {
		if i == 0 || v < ret {
			ret = v
		}
		i++
	}
	return ret
}

// Member checks if elem exists within the channel of values.
// The channel stops being read if the elem is found.
func Member[T comparable](ch <-chan T, elem T) bool {
	for v := range ch {
		if elem == v {
			return true
		}
	}
	return false
}

// Reduce reduces a channel using an accumulator.
func Reduce[S any, T any](ch <-chan S, acc T, f func(S, T) T) T {
	for v := range ch {
		acc = f(v, acc)
	}
	return acc
}

// Repeatedly returns a channel generated by calling `f` repeatedly.
// The channel does not close.
func Repeatedly[T any](f func() T) <-chan T {
	ch := make(chan T, defaultBufSize)

	go func() {
		for {
			ch <- f()
		}
	}()

	return ch
}

// Number has the "+" operator.
type Number interface {
	constraints.Integer | constraints.Unsigned | constraints.Float | constraints.Complex
}

// Sum sums up a channel of numbers.
func Sum[T Number](ch <-chan T) (ret T) {
	for v := range ch {
		ret += v
	}
	return ret
}

// Product multiples a channel of numbers together.
func Product[T Number](ch <-chan T) (ret T) {
	var i int
	for v := range ch {
		if i == 0 {
			ret = v
			i++
			continue
		}
		ret *= v
	}
	return ret
}
