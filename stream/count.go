package stream

// Count returns the count of items in the channel for which `f(item)` returns true.
func Count[T any](ch <-chan T, f FilterFunc[T]) int {
	var result int
	for v := range ch {
		if f(v) {
			result++
		}
	}
	return result
}

// CountWithIndex returns the count of items in the channel for which
// `f(index, item)` returns true.
func CountWithIndex[T any](ch <-chan T, f FilterFuncWithIndex[T]) int {
	var i, result int
	for v := range ch {
		if f(i, v) {
			result++
		}
		i++
	}
	return result
}
