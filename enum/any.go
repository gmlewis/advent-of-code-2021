package enum

// Any returns true if any f(item) call returns true.
func Any[T any](items []T, f func(T) bool) bool {
	for _, v := range items {
		if f(v) {
			return true
		}
	}
	return false
}

// AnyWithIndex returns true if any f(index, item) call returns true.
func AnyWithIndex[T any](items []T, f func(int, T) bool) bool {
	for i, v := range items {
		if f(i, v) {
			return true
		}
	}
	return false
}
