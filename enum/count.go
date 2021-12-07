package enum

// Count returns the count of items in the slice for which `f(item)` returns true.
func Count[T any](items []T, f FilterFunc[T]) int {
	var result int
	for _, v := range items {
		if f(v) {
			result++
		}
	}
	return result
}

// CountWithIndex returns the count of items in the slice for which
// `f(index, item)` returns true.
func CountWithIndex[T any](items []T, f FilterFuncWithIndex[T]) int {
	var result int
	for i, v := range items {
		if f(i, v) {
			result++
		}
	}
	return result
}
