package stream

// Find returns the first element for which f(value) returns true
// along with a boolean indicating a value was found.
func Find[T any](ch <-chan T, f FilterFunc[T]) (ret T, ok bool) {
	for v := range ch {
		if f(v) {
			return v, true
		}
	}
	return ret, false
}

// FindOr returns the first element for which f(value) returns true.
// If no element is found, defValue is returned.
func FindOr[T any](ch <-chan T, f FilterFunc[T], defValue T) T {
	for v := range ch {
		if f(v) {
			return v
		}
	}
	return defValue
}

// FindWithIndex returns the first element for which f(index, value) returns true
// along with a boolean indicating a value was found.
func FindWithIndex[T any](ch <-chan T, f FilterFuncWithIndex[T]) (ret T, ok bool) {
	var i int
	for v := range ch {
		if f(i, v) {
			return v, true
		}
		i++
	}
	return ret, false
}

// FindOrWithIndex returns the first element for which f(index, value) returns true.
// If no element is found, defValue is returned.
func FindOrWithIndex[T any](ch <-chan T, f FilterFuncWithIndex[T], defValue T) T {
	var i int
	for v := range ch {
		if f(i, v) {
			return v
		}
		i++
	}
	return defValue
}
