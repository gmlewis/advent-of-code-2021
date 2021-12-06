package stream

// Find returns the first element for which f(value) returns true.
// If no element is found, defValue is returned.
func Find[T any](ch <-chan T, f FilterFunc[T], defValue T) T {
	for v := range ch {
		if f(v) {
			return v
		}
	}
	return defValue
}

// FindWithIndex returns the first element for which f(index, value) returns true.
// If no element is found, defValue is returned.
func FindWithIndex[T any](ch <-chan T, f FilterFuncWithIndex[T], defValue T) T {
	var i int
	for v := range ch {
		if f(i, v) {
			return v
		}
		i++
	}
	return defValue
}
