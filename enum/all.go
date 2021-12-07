package enum

// All returns true if all f(item) calls return true.
func All[T any](items []T, f FilterFunc[T]) bool {
	for _, v := range items {
		if !f(v) {
			return false
		}
	}
	return true
}

// AllWithIndex returns true if all f(index, item) calls return true.
func AllWithIndex[T any](items []T, f FilterFuncWithIndex[T]) bool {
	for i, v := range items {
		if !f(i, v) {
			return false
		}
	}
	return true
}
