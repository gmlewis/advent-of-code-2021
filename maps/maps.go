// Package maps provides handy functional programming -style
// functions that operate on map data types.
package maps

import "constraints"

// Key returns the key of a map key/value pair.
func Key[K comparable, V any](key K, value V) K {
	return key
}

// Value returns the value of a map key/value pair.
func Value[K any, V any](key K, value V) V {
	return value
}

// ValueLen returns the length of the slice value of a map.
func ValueLen[K any, T any](key K, value []T) int {
	return len(value)
}

// All returns true if all f(key, value) calls return true.
func All[K comparable, V any](pairs map[K]V, f func(K, V) bool) bool {
	for k, v := range pairs {
		if !f(k, v) {
			return false
		}
	}
	return true
}

// Any returns true if any f(key, value) call returns true.
func Any[K comparable, V any](pairs map[K]V, f func(K, V) bool) bool {
	for k, v := range pairs {
		if f(k, v) {
			return true
		}
	}
	return false
}

// Count returns the count of items in the map for which `f(k, v)` returns true.
func Count[K comparable, V any](pairs map[K]V, f func(K, V) bool) int {
	var result int
	for k, v := range pairs {
		if f(k, v) {
			result++
		}
	}
	return result
}

// Map maps the (k,v) pairs to a slice of values.
func Map[K comparable, V any, T any](pairs map[K]V, f func(K, V) T) []T {
	result := []T{}
	for k, v := range pairs {
		result = append(result, f(k, v))
	}
	return result
}

// Reduce reduces a map using an accumulator.
func Reduce[K comparable, V any, T any](pairs map[K]V, acc T, f func(K, V, T) T) T {
	for k, v := range pairs {
		acc = f(k, v, acc)
	}
	return acc
}

// Number has the "+" operator.
type Number interface {
	constraints.Integer | constraints.Unsigned | constraints.Float | constraints.Complex
}

// SumKeys sums up the keys in a map.
func SumKeys[K Number, V any](pairs map[K]V) (ret K) {
	for k := range pairs {
		ret += k
	}
	return ret
}

// SumValues sums up the values in a map.
func SumValues[K comparable, V Number](pairs map[K]V) (ret V) {
	for _, v := range pairs {
		ret += v
	}
	return ret
}

// ProductKeys multiples the keys in a map together.
func ProductKeys[K Number, V any](pairs map[K]V) (ret K) {
	ret = K(1)
	for k := range pairs {
		ret *= k
	}
	return ret
}

// ProductValues multiples the values in a map together.
func ProductValues[K comparable, V Number](pairs map[K]V) (ret V) {
	ret = V(1)
	for _, v := range pairs {
		ret *= v
	}
	return ret
}
