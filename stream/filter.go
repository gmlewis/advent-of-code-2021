package stream

// FilterFunc takes a value and returns true if the
// value is to be kept.
type FilterFunc[T any] func(T) bool

// Filter filters a channel of values and keeps
// those for which f(value) returns true.
func Filter[T any](ch <-chan T, f FilterFunc[T]) []T {
	var result []T
	for v := range ch {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterMap filters a channel of mapped values and keeps
// those for which f(value) returns true.
func FilterMap[S any, T any](ch <-chan S, m func(S) T, f FilterFunc[T]) []S {
	var result []S
	for v := range ch {
		if f(m(v)) {
			result = append(result, v)
		}
	}
	return result
}

// FilterFuncWithIndex takes an index and value and returns true if the
// value is to be kept.
type FilterFuncWithIndex[T any] func(int, T) bool

// FilterWithIndex filters a channel of values and keeps
// those for which f(index, value) returns true.
func FilterWithIndex[T any](ch <-chan T, f FilterFuncWithIndex[T]) []T {
	var result []T
	var i int
	for v := range ch {
		if f(i, v) {
			result = append(result, v)
		}
		i++
	}
	return result
}
