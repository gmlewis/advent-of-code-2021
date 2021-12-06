package stream

// Uniq creates a stream that only emits elements if they are unique.
//
// Keep in mind that, in order to know if an element is unique or not, this
// function needs to store all unique values emitted by the stream. Therefore, if
// the stream is infinite, the number of elements stored will grow infinitely,
// never being garbage-collected.
func Uniq[T comparable](ch <-chan T) <-chan T {
	outCh := make(chan T, defaultBufSize)

	go func() {
		seen := map[T]struct{}{}
		for v := range ch {
			if _, ok := seen[v]; ok {
				continue
			}
			outCh <- v
			seen[v] = struct{}{}
		}

		close(outCh)
	}()

	return outCh
}
