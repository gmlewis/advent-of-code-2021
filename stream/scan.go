package stream

// Scan creates a channel that applies the given function to each element,
// emits the result and uses the same result as the accumulator for the
// next computation. It uses the given acc as the starting value.
func Scan[T any](ch <-chan T, acc T, f func(a, b T) T) <-chan T {
	outCh := make(chan T, defaultBufSize)

	go func() {
		for v := range ch {
			nv := f(acc, v)
			outCh <- nv
			acc = nv
		}

		close(outCh)
	}()

	return outCh
}
