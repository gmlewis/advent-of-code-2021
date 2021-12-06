package enum

// FilterFunc takes a value and returns true if the
// value is to be kept.
type FilterFunc[T any] func(T) bool

// Filter filters a slice of values and keeps
// those for which f(value) returns true.
func Filter[T any](values []T, f FilterFunc[T]) []T {
	var result []T
	for _, v := range values {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterMap filters a slice of mapped values and keeps
// those for which f(value) returns true.
func FilterMap[S any, T any](values []S, m func(S) T, f FilterFunc[T]) []S {
	var result []S
	for _, v := range values {
		if f(m(v)) {
			result = append(result, v)
		}
	}
	return result
}

// FilterFuncWithIndex takes an index and value and returns true if the
// value is to be kept.
type FilterFuncWithIndex[T any] func(int, T) bool

// FilterWithIndex filters a slice of values and keeps
// those for which f(index, value) returns true.
func FilterWithIndex[T any](values []T, f FilterFuncWithIndex[T]) []T {
	var result []T
	for i, v := range values {
		if f(i, v) {
			result = append(result, v)
		}
	}
	return result
}
