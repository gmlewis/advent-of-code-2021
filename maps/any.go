package maps

// Any returns true if any f(key, value) call returns true.
func Any[K comparable, V any](pairs map[K]V, f func(K, V) bool) bool {
	for k, v := range pairs {
		if f(k, v) {
			return true
		}
	}
	return false
}
