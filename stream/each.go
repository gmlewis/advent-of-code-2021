package stream

// Each processes each item with the provided function.
func Each[T any](ch <-chan T, f func(item T)) {
	for item := range ch {
		f(item)
	}
}

// EachWithIndex iterates over a channel and calls the provided
// function with its index and value.
func EachWithIndex[T any](ch <-chan T, f func(i int, value T)) {
	var i int
	for value := range ch {
		f(i, value)
		i++
	}
}
