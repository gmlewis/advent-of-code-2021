package maps

// All returns true if all f(key, value) calls return true.
func All[K comparable, V any](pairs map[K]V, f func(K, V) bool) bool {
	for k, v := range pairs {
		if !f(k, v) {
			return false
		}
	}
	return true
}
