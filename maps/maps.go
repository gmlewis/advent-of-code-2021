// Package maps provides handy functional programming -style
// functions that operate on map data types.
package maps

// Any returns true if any f(key,value) returns true.
func Any[K comparable, V any](pairs map[K]V, f func(K, V) bool) bool {
	for k, v := range pairs {
		if f(k, v) {
			return true
		}
	}
	return false
}

// Reduce reduces a map using an accumulator.
func Reduce[K comparable, V any, T any](pairs map[K]V, acc T, f func(K, V, T) T) T {
	for k, v := range pairs {
		acc = f(k, v, acc)
	}
	return acc
}
