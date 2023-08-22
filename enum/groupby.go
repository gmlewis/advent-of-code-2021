package enum

// GroupBy splits the items into groups based on keyFunc and valueFunc.
func GroupBy[K comparable, V any, T any](items []T, keyFunc func(T) K, valueFunc func(T) V) map[K][]V {
	result := map[K][]V{}
	for _, item := range items {
		k := keyFunc(item)
		result[k] = append(result[k], valueFunc(item))
	}
	return result
}
