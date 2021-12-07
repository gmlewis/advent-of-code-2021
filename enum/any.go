package enum

// Any returns true if any f(item) call returns true.
func Any[T any](items []T, f FilterFunc[T]) bool {
	for _, v := range items {
		if f(v) {
			return true
		}
	}
	return false
}

// AnyWithIndex returns true if any f(index, item) call returns true.
func AnyWithIndex[T any](items []T, f FilterFuncWithIndex[T]) bool {
	for i, v := range items {
		if f(i, v) {
			return true
		}
	}
	return false
}
