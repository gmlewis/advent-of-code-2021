package stream

// Repeatedly returns a channel generated by calling `f` repeatedly.
// The channel does not close.
func Repeatedly[T any](f func() T) <-chan T {
	ch := make(chan T, defaultBufSize)

	go func() {
		for {
			ch <- f()
		}
	}()

	return ch
}
