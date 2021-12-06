package enum

// Find returns the first element for which f(value) returns true
// along with a boolean indicating a value was found.
func Find[T any](values []T, f FilterFunc[T]) (ret T, ok bool) {
	for _, v := range values {
		if f(v) {
			return v, true
		}
	}
	return ret, false
}

// FindOr returns the first element for which f(value) returns true.
// If no element is found, defValue is returned.
func FindOr[T any](values []T, f FilterFunc[T], defValue T) T {
	for _, v := range values {
		if f(v) {
			return v
		}
	}
	return defValue
}

// FindWithIndex returns the first element for which f(index, value) returns true
// along with a boolean indicating a value was found.
func FindWithIndex[T any](values []T, f FilterFuncWithIndex[T]) (ret T, ok bool) {
	for i, v := range values {
		if f(i, v) {
			return v, true
		}
	}
	return ret, false
}

// FindOrWithIndex returns the first element for which f(index, value) returns true.
// If no element is found, defValue is returned.
func FindOrWithIndex[T any](values []T, f FilterFuncWithIndex[T], defValue T) T {
	for i, v := range values {
		if f(i, v) {
			return v
		}
	}
	return defValue
}
