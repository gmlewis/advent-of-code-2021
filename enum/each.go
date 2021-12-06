package enum

// Each processes each item with the provided function.
func Each[T any](items []T, f func(item T)) {
	for _, item := range items {
		f(item)
	}
}

// EachWithIndex iterates over a slice and calls the provided
// function with its index and value.
func EachWithIndex[T any](items []T, f func(i int, value T)) {
	for i, value := range items {
		f(i, value)
	}
}
