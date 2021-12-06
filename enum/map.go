package enum

// Map maps a slice of values from one type to another
// using the provided func f.
func Map[S any, T any](values []T, f func(a T) S) []S {
	result := make([]S, 0, len(values))
	for _, v := range values {
		result = append(result, f(v))
	}
	return result
}
