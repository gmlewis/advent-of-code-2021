package stream

// Reduce reduces a channel using an accumulator.
func Reduce[S any, T any](ch <-chan S, acc T, f func(S, T) T) T {
	for v := range ch {
		acc = f(v, acc)
	}
	return acc
}
