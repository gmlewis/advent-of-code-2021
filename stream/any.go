package stream

// Any returns true if any f(item) call returns true.
func Any[T any](ch <-chan T, f func(T) bool) bool {
	for v := range ch {
		if f(v) {
			return true
		}
	}
	return false
}

// AnyWithIndex returns true if any f(index, item) call returns true.
func AnyWithIndex[T any](ch <-chan T, f func(int, T) bool) bool {
	var i int
	for v := range ch {
		if f(i, v) {
			return true
		}
		i++
	}
	return false
}
