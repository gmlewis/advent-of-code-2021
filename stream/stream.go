// Package stream provides handy functional programming -style
// functions that operate on enumerable data types using channels.
package stream

import (
	"golang.org/x/exp/constraints"
)

const (
	defaultBufSize = 1000
)

// First returns the first item of the provided chan or its zero value.
func First[T any](ch <-chan T) (ret T) {
	if ch == nil {
		return ret
	}
	return <-ch
}

// Length returns the length of the provided channel.
func Length[T any](ch <-chan T) int {
	if ch == nil {
		return 0
	}
	var i int
	for _ = range ch {
		i++
	}
	return i
}

// Take returns the first n items from the provided chan.
func Take[T any](ch <-chan T, n int) []T {
	if ch == nil {
		return nil
	}
	ret := make([]T, 0, n)
	for i := 0; i < n; i++ {
		v, ok := <-ch
		if !ok {
			break
		}
		ret = append(ret, v)
	}
	return ret
}

// Number has the "+" operator.
type Number interface {
	constraints.Integer | constraints.Unsigned | constraints.Float | constraints.Complex
}

// Sum sums up a channel of numbers or the zero value if the channel is empty.
func Sum[T Number](ch <-chan T) (ret T) {
	if ch == nil {
		return ret
	}
	for v := range ch {
		ret += v
	}
	return ret
}

// Product multiples a channel of numbers together or the
// zero value if the channel is empty.
func Product[T Number](ch <-chan T) (ret T) {
	if ch == nil {
		return ret
	}
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
