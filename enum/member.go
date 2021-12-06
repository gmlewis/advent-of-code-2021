package enum

// Member checks if elem exists within values.
func Member[T comparable](values []T, elem T) bool {
	for _, v := range values {
		if elem == v {
			return true
		}
	}
	return false
}
