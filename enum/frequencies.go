package enum

// Frequencies returns a map with keys as unique elements of the
// provided items and the values as the count of every item.
func Frequencies[T comparable](items []T) map[T]int {
	result := map[T]int{}
	for _, item := range items {
		result[item]++
	}
	return result
}

// FrequenciesBy returns a map with keys as unique elements of
// keyFunc(item) and the values as the count of every item.
func FrequenciesBy[S any, T comparable](items []S, keyFunc func(S) T) map[T]int {
	result := map[T]int{}
	for _, item := range items {
		result[keyFunc(item)]++
	}
	return result
}
