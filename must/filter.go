package must

// FilterFunc takes a value and returns true if the
// value is to be kept.
type FilterFunc[T any] func(T) bool

// Filter filters a slice of values and keeps
// those for which f returns true.
func Filter[T any](values []T, f FilterFunc[T]) []T {
	var result []T
	for _, v := range values {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
