package maps

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
