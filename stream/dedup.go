package stream

// Dedup returns a channel where all consecutive duplicated
// elements are collapsed to a single element.
func Dedup[T comparable](ch <-chan T) <-chan T {
	outCh := make(chan T, defaultBufSize)

	go func() {
		var i int
		var last T
		for v := range ch {
			if i == 0 || v != last {
				outCh <- v
			}
			last = v
			i++
		}

		close(outCh)
	}()

	return outCh
}
