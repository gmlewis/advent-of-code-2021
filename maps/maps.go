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

// Map maps the (k,v) pairs to a slice of values.
func Map[K comparable, V any, T any](pairs map[K]V, f func(K, V) T) []T {
	result := []T{}
	for k, v := range pairs {
		result = append(result, f(k, v))
	}
	return result
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
