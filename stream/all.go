package stream

// All returns true if all f(item) calls return true.
func All[T any](ch <-chan T, f func(T) bool) bool {
	for v := range ch {
		if !f(v) {
			return false
		}
	}
	return true
}

// AllWithIndex returns true if all f(index, item) calls return true.
func AllWithIndex[T any](ch <-chan T, f func(int, T) bool) bool {
	var i int
	for v := range ch {
		if !f(i, v) {
			return false
		}
		i++
	}
	return true
}
