package stream

// Reduce reduces a channel using an accumulator and `f(item, acc)`.
func Reduce[S any, T any](ch <-chan S, acc T, f func(S, T) T) T {
	for v := range ch {
		acc = f(v, acc)
	}
	return acc
}

// ReduceWithIndex reduces a slice using an accumulator and `f(index, item, acc)`.
func ReduceWithIndex[S any, T any](ch <-chan S, acc T, f func(int, S, T) T) T {
	var i int
	for v := range ch {
		acc = f(i, v, acc)
		i++
	}
	return acc
}
