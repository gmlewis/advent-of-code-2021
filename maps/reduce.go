package maps

// Reduce reduces a map using an accumulator.
func Reduce[K comparable, V any, T any](pairs map[K]V, acc T, f func(K, V, T) T) T {
	for k, v := range pairs {
		acc = f(k, v, acc)
	}
	return acc
}
