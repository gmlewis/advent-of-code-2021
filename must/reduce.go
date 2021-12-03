package must

// Reduce reduces slices using an accumulator.
func Reduce[S any, T any](values []S, acc T, f func(S, T) T) T {
	for _, val := range values {
		acc = f(val, acc)
	}
	return acc
}
