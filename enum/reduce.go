package enum

// Reduce reduces a slice using an accumulator and `f(item, acc)`.
func Reduce[S any, T any](items []S, acc T, f func(S, T) T) T {
	for _, v := range items {
		acc = f(v, acc)
	}
	return acc
}

// ReduceWithIndex reduces a slice using an accumulator and `f(index, item, acc)`.
func ReduceWithIndex[S any, T any](items []S, acc T, f func(int, S, T) T) T {
	for i, v := range items {
		acc = f(i, v, acc)
	}
	return acc
}
