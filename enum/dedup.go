package enum

// Dedup returns a slice where all consecutive duplicated
// elements are collapsed to a single element.
func Dedup[T comparable](values []T) []T {
	result := []T{}
	for i, v := range values {
		if i == 0 || v != values[i-1] {
			result = append(result, v)
		}
	}
	return result
}
