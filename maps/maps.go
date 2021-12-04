// Package maps provides handy functional programming -style
// functions that operate on map data types.
package maps

// Reduce reduces a map using an accumulator.
func Reduce[K comparable, V any, T any](values map[K]V, acc T, f func(K, V, T) T) T {
	for k, v := range values {
		acc = f(k, v, acc)
	}
	return acc
}
