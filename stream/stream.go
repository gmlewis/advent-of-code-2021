// Package stream provides handy functional programming -style
// functions that operate on enumerable data types using channels.
package stream

import (
	"constraints"
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
