package stream

// Any returns true if any f(item) call returns true.
func Any[T any](ch <-chan T, f FilterFunc[T]) bool {
	for v := range ch {
		if f(v) {
			return true
		}
	}
	return false
}

// AnyWithIndex returns true if any f(index, item) call returns true.
func AnyWithIndex[T any](ch <-chan T, f FilterFuncWithIndex[T]) bool {
	var i int
	for v := range ch {
		if f(i, v) {
			return true
		}
		i++
	}
	return false
}
