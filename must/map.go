//go:build go1.18

package must

// Map maps a slice of values from one type to another
// using the provided func f.
func Map[S any, T any](slice []T, f func(a T) S) []S {
	result := make([]S, 0, len(slice))
	for _, val := range slice {
		result = append(result, f(val))
	}
	return result
}
