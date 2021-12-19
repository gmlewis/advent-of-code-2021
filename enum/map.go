package enum

// Map maps a slice of values from one type to another
// using the provided func f.
func Map[S any, T any](values []T, f func(value T) S) []S {
	result := make([]S, 0, len(values))
	for _, v := range values {
		result = append(result, f(v))
	}
	return result
}

// MapWithIndex maps a slice of values from one type to another
// using the provided func f.
func MapWithIndex[S any, T any](values []T, f func(index int, value T) S) []S {
	result := make([]S, 0, len(values))
	for i, v := range values {
		result = append(result, f(i, v))
	}
	return result
}
